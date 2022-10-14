package bootstrap

import (
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/getsentry/sentry-go"
)

func InitSentry(cfg *confpb.Sentry) error {
	if cfg != nil {
		return sentry.Init(sentry.ClientOptions{
			Dsn:              cfg.GetDsn(),
			AttachStacktrace: cfg.GetAttachStackTrace(), // recommended
			Release:          cfg.GetRelease(),
			Environment:      cfg.GetEnvironment(),
			ServerName:       cfg.GetServerName(),
		})
	}
	return nil
}
