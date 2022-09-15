package zapx

import (
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

func TestZapLogger(t *testing.T) {
	//encoder := zapcore.EncoderConfig{
	//	TimeKey:        "t",
	//	LevelKey:       "level",
	//	NameKey:        "logger",
	//	CallerKey:      "caller",
	//	MessageKey:     "msg",
	//	StacktraceKey:  "stack",
	//	EncodeTime:     zapcore.ISO8601TimeEncoder,
	//	LineEnding:     zapcore.DefaultLineEnding,
	//	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	//	EncodeDuration: zapcore.SecondsDurationEncoder,
	//	EncodeCaller:   zapcore.FullCallerEncoder,
	//}
	logger := NewZapLogger()
	zlog := log.NewHelper(logger)
	zlog.Infow("name", "kratos", "from", "opensource")
	zlog.Infow("name", "kratos", "from")

	// zap stdout/stderr Sync bugs in OSX, see https://github.com/uber-go/zap/issues/370
	_ = logger.Sync()
}

func TestNewZapLoggerWithConfig(t *testing.T) {
	logpath := "logs"
	type args struct {
		cfg      *confpb.ZapLog
		logLevel log.Level
		kvs      []interface{}
		debug    []bool
	}
	tests := []struct {
		name string
		args args
		want *ZapLogger
	}{
		{
			args: args{
				cfg: &confpb.ZapLog{
					//Kafka: nil,
					File: &confpb.FileHook{
						Path:         logpath,
						Filename:     "test",
						MaxAge:       nil,
						RotationTime: nil,
						Pattern:      "",
					},
				},
				logLevel: log.LevelInfo,
				kvs:      []interface{}{"case", "json format", "msg", "1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewZapLoggerWithConfig(tt.args.cfg)
			_ = logger.Log(tt.args.logLevel, tt.args.kvs...)
		})
	}
}
