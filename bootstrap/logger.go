package bootstrap

import (
	"github.com/aesoper101/kratos-utils/internal/log/aliyunx"
	"github.com/aesoper101/kratos-utils/internal/log/fluentx"
	"github.com/aesoper101/kratos-utils/internal/log/logrusx"
	"github.com/aesoper101/kratos-utils/internal/log/zapx"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"os"
)

func NewLoggerProvide(cfg *confpb.Log, serviceInfo *ServiceInfo) log.Logger {
	var logger log.Logger

	if cfg != nil {
		if cfg.GetZap() != nil {
			logger = zapx.NewZapLoggerWithConfig(cfg.GetZap(), cfg.GetDebug())
		} else if cfg.GetLogrus() != nil {
			logger = logrusx.NewLogrusLoggerWithConfig(cfg.GetLogrus(), cfg.GetDebug())
		} else if cfg.GetAliyun() != nil {
			logger = aliyunx.NewAliyunLoggerWithConfig(cfg.GetAliyun())
		} else if cfg.GetFluent() != nil {
			logger = fluentx.NewFluentLoggerWithConfig(cfg.GetFluent())
		} else {
			logger = log.NewStdLogger(os.Stdout)
		}
	} else {
		logger = log.NewStdLogger(os.Stdout)
	}

	return log.With(
		logger,
		"service.id", serviceInfo.Id,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}
