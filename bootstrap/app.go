package bootstrap

import (
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
)

type App struct {
	config  *AppConfig
	srvInfo *ServiceInfo
	app     *kratos.App
}

type AppConfig struct {
	Opensergo *confpb.OpenSergo
	Registry  *confpb.Registry
	Tracer    *confpb.Tracer
	Sentry    *confpb.Sentry
	Logger    log.Logger
}

var (
	ProviderSet = wire.NewSet(wire.Struct(new(AppConfig), "*"), NewLoggerProvide, NewApp)
)

func NewApp(serviceInfo *ServiceInfo, cfg *AppConfig, servers ...transport.Server) *App {
	kratosApp := kratos.New(
		kratos.ID(serviceInfo.GetInstanceId()),
		kratos.Name(serviceInfo.Name),
		kratos.Version(serviceInfo.Version),
		kratos.Metadata(serviceInfo.Metadata),
		kratos.Logger(cfg.Logger),
		kratos.Server(
			servers...,
		),
		kratos.Registrar(NewRegistrarProvider(cfg.Registry)),
	)

	return &App{app: kratosApp, config: cfg, srvInfo: serviceInfo}
}

func (app *App) Run() error {
	if app.config.Tracer != nil {
		if _, err := NewTracerProvider(app.config.Tracer, app.srvInfo); err != nil {
			return err
		}
	}

	if app.config.Opensergo != nil {
		if err := InitOpenSergo(app.app, app.config.Opensergo); err != nil {
			return err
		}
	}

	if app.config.Sentry != nil {
		if err := InitSentry(app.config.Sentry); err != nil {
			return err
		}
	}

	return app.app.Run()
}
