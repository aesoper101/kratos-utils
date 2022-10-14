package fluentx

import (
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/contrib/log/fluent/v2"
)

func NewFluentLoggerWithConfig(cfg *confpb.FluentLog) *fluent.Logger {
	addr := "tcp://127.0.0.1:24224"
	if cfg == nil {
		return nil
	}

	if cfg.GetAddr() != "" {
		addr = cfg.GetAddr()
	}

	var opts []fluent.Option
	if cfg.MaxRetry > 0 {
		opts = append(opts, fluent.WithMaxRetry(int(cfg.MaxRetryWait)))
	}
	if cfg.RetryWait > 0 {
		opts = append(opts, fluent.WithMaxRetryWait(int(cfg.RetryWait)))
	}
	if cfg.BufferLimit > 0 {
		opts = append(opts, fluent.WithBufferLimit(int(cfg.BufferLimit)))
	}
	if cfg.GetWriteTimeout() != nil {
		opts = append(opts, fluent.WithTimeout(cfg.GetWriteTimeout().AsDuration()))
	}
	if len(cfg.GetTagPrefix()) > 0 {
		opts = append(opts, fluent.WithTagPrefix(cfg.GetTagPrefix()))
	}
	opts = append(opts, fluent.WithAsync(cfg.GetAsync()), fluent.WithForceStopAsyncSend(cfg.GetForceStopAsyncSend()))
	logger, _ := fluent.NewLogger(addr)

	return logger
}
