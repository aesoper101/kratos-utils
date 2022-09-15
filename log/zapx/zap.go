package zapx

import (
	"fmt"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var _ log.Logger = (*ZapLogger)(nil)

// ZapLogger is a logger impl.
type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

type zapConfig struct {
	encoder zapcore.Encoder
	level   zap.AtomicLevel
	opts    []zap.Option
	cores   []zapcore.Core
}

func NewZapLoggerWithConfig(cfg *confpb.ZapLog, debug ...bool) *ZapLogger {
	var options []Option

	encoderConfig := zap.NewProductionEncoderConfig()
	atomicLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	if len(debug) > 0 && debug[0] {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		atomicLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	if cfg != nil {
		if strings.ToLower(cfg.GetFormatter()) == "text" {
			encoder = zapcore.NewConsoleEncoder(encoderConfig)
		}

		if cfg.GetFile() != nil {
			options = append(options, WithSyncHook(zapcore.NewCore(encoder, zapcore.AddSync(newFileHook(cfg.GetFile())), atomicLevel)))
		}

		if cfg.GetKafka() != nil {
			options = append(options, WithSyncHook(zapcore.NewCore(encoder, zapcore.AddSync(newKafkaHook(cfg.GetKafka())), atomicLevel)))
		}
	}

	options = append(options, WithEncoder(encoder), WithLevel(atomicLevel))

	return NewZapLogger(options...)
}

// NewZapLogger return a zap logger.
func NewZapLogger(opts ...Option) *ZapLogger {
	cfg := &zapConfig{
		encoder: zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		level:   zap.NewAtomicLevelAt(zap.InfoLevel),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	cores := []zapcore.Core{
		zapcore.NewCore(cfg.encoder, zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
		), cfg.level),
	}
	cores = append(cores, cfg.cores...)

	core := zapcore.NewTee(cores...)

	logger := zap.New(core, cfg.opts...)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	go func(logger *zap.Logger) {
		<-c
		_ = logger.Sync()
	}(logger)

	return &ZapLogger{log: logger, Sync: logger.Sync}
}

// Log Implementation of logger interface.
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}

type Option func(logger *zapConfig)

func WithEncoder(encoder zapcore.Encoder) Option {
	return func(logger *zapConfig) {
		logger.encoder = encoder
	}
}

func WithLevel(level zap.AtomicLevel) Option {
	return func(logger *zapConfig) {
		logger.level = level
	}
}

func WithZapOption(opt zap.Option) Option {
	return func(logger *zapConfig) {
		logger.opts = append(logger.opts, opt)
	}
}

func WithSyncHook(core zapcore.Core) Option {
	return func(logger *zapConfig) {
		logger.cores = append(logger.cores, core)
	}
}
