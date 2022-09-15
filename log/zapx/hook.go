package zapx

import (
	"crypto/tls"
	"github.com/Shopify/sarama"
	"github.com/aesoper101/go-utils/str"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lestrrat-go/strftime"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"path/filepath"
	"time"
)

func newFileHook(cfg *confpb.FileHook) io.Writer {
	path := cfg.GetPath()
	if path == "" {
		path = "./"
	}
	filename := str.ToSnake(cfg.GetFilename())
	if filename == "" {
		path = "app"
	}
	rotationTime := time.Hour
	if cfg.GetRotationTime() != nil {
		rotationTime = cfg.GetRotationTime().AsDuration()
	}
	maxAge := time.Duration(604800) * time.Second
	if cfg.GetMaxAge() != nil {
		rotationTime = cfg.GetMaxAge().AsDuration()
	}

	pattern := "%Y%m%d"
	if cfg.GetPattern() != "" {
		if _, err := strftime.New(cfg.GetPattern()); err == nil {
			pattern = cfg.GetPattern()
		}
	}

	writer, _ := rotatelogs.New(
		filepath.Join(path, filename+".log."+pattern),
		rotatelogs.WithLinkName(filepath.Join(path, filename+".log")),
		rotatelogs.WithMaxAge(rotationTime),
		rotatelogs.WithRotationTime(maxAge),
	)
	return writer
}

func newKafkaHook(cfg *confpb.KafkaHook) io.Writer {
	if cfg != nil {
		return nil
	}

	logId, defaultTopic := "kafka_log", "topic_log"
	if cfg.Id != "" {
		logId = cfg.Id
	}
	if cfg.GetDefaultTopic() != "" {
		defaultTopic = cfg.GetDefaultTopic()
	}

	hook, err := NewKafkaLogrusHook(
		logId,
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		cfg.Brokers,
		defaultTopic,
		nil,
	)
	if err != nil {
		return nil
	}

	return hook
}

type KafkaZapHook struct {
	id           string
	defaultTopic string
	levels       []logrus.Level
	producer     sarama.AsyncProducer
}

// NewKafkaLogrusHook creates a new KafkaHook
func NewKafkaLogrusHook(id string,
	levels []logrus.Level,
	brokers []string,
	defaultTopic string,
	tls *tls.Config) (*KafkaZapHook, error) {
	var err error
	var producer sarama.AsyncProducer
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	kafkaConfig.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	kafkaConfig.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	// check here if provided *tls.Config is not nil and assign to the sarama config
	// NOTE: we automatically enabled the TLS config because sarama would error out if our
	//       config were non-nil but disabled. To avoid issue further down the stack, we enable.
	if tls != nil {
		kafkaConfig.Net.TLS.Enable = true
		kafkaConfig.Net.TLS.Config = tls
	}

	if producer, err = sarama.NewAsyncProducer(brokers, kafkaConfig); err != nil {
		return nil, err
	}

	go func() {
		for err := range producer.Errors() {
			log.Printf("Failed to send log entry to Kafka: %v\n", err)
		}
	}()

	hook := &KafkaZapHook{
		id,
		defaultTopic,
		levels,
		//formatter,
		producer,
	}

	return hook, nil
}

func (k *KafkaZapHook) Write(b []byte) (int, error) {
	var partitionKey sarama.ByteEncoder
	partitionKey = sarama.ByteEncoder(b)

	value := sarama.ByteEncoder(b)

	topic := k.defaultTopic

	k.producer.Input() <- &sarama.ProducerMessage{
		Key:   partitionKey,
		Topic: topic,
		Value: value,
	}
	return len(value), nil
}
