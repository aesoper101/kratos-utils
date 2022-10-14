package logrusx

import (
	"bytes"
	"github.com/aesoper101/go-utils/filex"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestLoggerLog(t *testing.T) {
	tests := map[string]struct {
		level     logrus.Level
		formatter logrus.Formatter
		logLevel  log.Level
		kvs       []interface{}
		want      string
	}{
		"json format": {
			level:     logrus.InfoLevel,
			formatter: &logrus.JSONFormatter{},
			logLevel:  log.LevelInfo,
			kvs:       []interface{}{"case", "json format", "msg", "1"},
			want:      `{"case":"json format","level":"info","msg":"1"`,
		},
		"level unmatch": {
			level:     logrus.InfoLevel,
			formatter: &logrus.JSONFormatter{},
			logLevel:  log.LevelDebug,
			kvs:       []interface{}{"case", "level unmatch", "msg", "1"},
			want:      "",
		},
		"no tags": {
			level:     logrus.InfoLevel,
			formatter: &logrus.JSONFormatter{},
			logLevel:  log.LevelInfo,
			kvs:       []interface{}{"msg", "1"},
			want:      `{"level":"info","msg":"1"`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			output := new(bytes.Buffer)
			logger := NewLogrusLogger(Level(test.level), Formatter(test.formatter), Output(output))
			_ = logger.Log(test.logLevel, test.kvs...)
			if !strings.HasPrefix(output.String(), test.want) {
				t.Errorf("strings.HasPrefix(output.String(), test.want) got %v want: %v", strings.HasPrefix(output.String(), test.want), true)
			}
		})
	}
}

func TestNewLogrusLoggerWithConfig(t *testing.T) {
	logpath := "logs"
	type args struct {
		cfg      *confpb.LogrusLog
		logLevel log.Level
		kvs      []interface{}
	}
	tests := []struct {
		name string
		args args
		want log.Logger
	}{
		{
			name: "file hook",
			args: args{cfg: &confpb.LogrusLog{
				//Formatter: "text",
				Kafka: nil,
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
			logger := NewLogrusLoggerWithConfig(tt.args.cfg)
			_ = logger.Log(tt.args.logLevel, tt.args.kvs...)
			if !filex.IsExists(tt.args.cfg.File.Path) {
				t.Fatal("log failed")
			}
		})
	}

}
