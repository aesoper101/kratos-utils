package logrusx

import (
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var _ log.Logger = (*LogrusLogger)(nil)

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLoggerWithConfig(cfg *confpb.LogrusLog, debug ...bool) *LogrusLogger {
	var options []Option
	if cfg != nil {
		if cfg.File != nil {
			options = append(options, newLfShook(cfg.File, cfg.Formatter))
		}
		if cfg.Kafka != nil {
			options = append(options, newKafkaHook(cfg.Kafka, cfg.Formatter))
		}
	}
	if len(debug) > 0 && debug[0] {
		options = append(options, Level(logrus.DebugLevel))
	}
	options = append(options, Formatter(getFormatter(cfg.Formatter)))
	return NewLogrusLogger(options...)
}

func getFormatter(formatter string) logrus.Formatter {
	if strings.ToLower(formatter) == "text" {
		return &logrus.TextFormatter{}
	}

	return &logrus.JSONFormatter{}
}

func NewLogrusLogger(options ...Option) *LogrusLogger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	for _, option := range options {
		option(logger)
	}

	return &LogrusLogger{
		log: logger,
	}
}

func (l *LogrusLogger) Log(level log.Level, keyvals ...interface{}) (err error) {
	var (
		logrusLevel logrus.Level
		fields      logrus.Fields = make(map[string]interface{})
		msg         string
	)

	switch level {
	case log.LevelDebug:
		logrusLevel = logrus.DebugLevel
	case log.LevelInfo:
		logrusLevel = logrus.InfoLevel
	case log.LevelWarn:
		logrusLevel = logrus.WarnLevel
	case log.LevelError:
		logrusLevel = logrus.ErrorLevel
	default:
		logrusLevel = logrus.DebugLevel
	}

	if logrusLevel > l.log.Level {
		return
	}

	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "")
	}
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			continue
		}
		if key == logrus.FieldKeyMsg {
			msg, _ = keyvals[i+1].(string)
			continue
		}
		fields[key] = keyvals[i+1]
	}

	if len(fields) > 0 {
		l.log.WithFields(fields).Log(logrusLevel, msg)
	} else {
		l.log.Log(logrusLevel, msg)
	}

	return
}

type Option func(log *logrus.Logger)

func Level(level logrus.Level) Option {
	return func(log *logrus.Logger) {
		log.Level = level
	}
}

func Output(w io.Writer) Option {
	return func(log *logrus.Logger) {
		log.Out = w
	}
}

func Formatter(formatter logrus.Formatter) Option {
	return func(log *logrus.Logger) {
		log.Formatter = formatter
	}
}

func AddHook(hook logrus.Hook) Option {
	return func(log *logrus.Logger) {
		if hook != nil {
			log.AddHook(hook)
		}
	}
}
