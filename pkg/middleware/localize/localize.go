package localize

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

type (
	localizerKey struct{}
)

func I18N(bundle *I18nBundle) middleware.Middleware {
	if bundle == nil {
		panic("i18n bundle must be set.")
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				accept := tr.RequestHeader().Get("accept-language")
				localizer := NewLocalizer(bundle, accept)

				ctx = context.WithValue(ctx, localizerKey{}, localizer)
			}
			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) *Localizer {
	return ctx.Value(localizerKey{}).(*Localizer)
}
