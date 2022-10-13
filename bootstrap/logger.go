package bootstrap

import (
	"github.com/aesoper101/kratos-utils/log/aliyunx"
	"github.com/aesoper101/kratos-utils/log/fluentx"
	"github.com/aesoper101/kratos-utils/log/logrusx"
	"github.com/aesoper101/kratos-utils/log/zapx"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"os"
)

func NewLoggerProvide(cfg *confpb.Log, serviceInfo *ServiceInfo, debug ...bool) log.Logger {
	d := false
	if len(debug) > 0 {
		d = debug[0]
	}

	var logger log.Logger

	if cfg != nil {
		if cfg.GetZap() != nil {
			logger = zapx.NewZapLoggerWithConfig(cfg.GetZap(), d)
		} else if cfg.GetLogrus() != nil {
			logger = logrusx.NewLogrusLoggerWithConfig(cfg.GetLogrus(), d)
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
