package logrusx

import (
	"github.com/aesoper101/go-utils/str"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lestrrat-go/strftime"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	kl "github.com/tracer0tong/kafkalogrus"
	"path/filepath"
	"time"
)

func newLfShook(cfg *confpb.FileHook, formatter string) Option {
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

	return AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
		},
		getFormatter(formatter),
	))
}

func newKafkaHook(cfg *confpb.KafkaHook, formatter string) Option {
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

	hook, err := kl.NewKafkaLogrusHook(
		logId,
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		getFormatter(formatter),
		cfg.Brokers,
		defaultTopic,
		false,
		nil,
	)
	if err != nil {
		return nil
	}

	return AddHook(hook)
}
