package bootstrap

import (
	"errors"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func NewTracerProvider(cfg *confpb.Tracer, serviceInfo *ServiceInfo) error {
	if cfg == nil {
		return errors.New("tracer config can not be nil")
	}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.Endpoint)))
	if err != nil {
		return err
	}

	tp := traceSdk.NewTracerProvider(
		traceSdk.WithSampler(traceSdk.ParentBased(traceSdk.TraceIDRatioBased(1.0))),
		traceSdk.WithBatcher(exp),
		traceSdk.WithResource(resource.NewSchemaless(
			semConv.ServiceNameKey.String(serviceInfo.Name),
			semConv.ServiceVersionKey.String(serviceInfo.Version),
			semConv.ServiceInstanceIDKey.String(serviceInfo.Id),
			attribute.String("env", cfg.Env),
		)),
	)

	otel.SetTracerProvider(tp)

	return nil
}
