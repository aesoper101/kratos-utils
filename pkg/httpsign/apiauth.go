package httpsign

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"time"
)

func Handler(opts ...Option) middleware.Middleware {
	o := &options{
		appIdGetter:        defaultAppIdGetter,
		signatureGetter:    defaultSignatureGetter,
		timestampGetter:    defaultTimestampGetter,
		timeout:            300, // 300 seconds
		signatureValidator: defaultSignatureValidator,
	}

	for _, opt := range opts {
		opt(o)
	}

	if o.secretGetter != nil {
		panic("secretGetter must be set.")
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if _, ok := transport.FromServerContext(ctx); ok {
				appId := o.appIdGetter(ctx)
				if len(appId) == 0 {
					return nil, errors.Forbidden("ERROR_MISSING_APPID", "missing query parameter: appid")
				}

				appSecret := o.secretGetter(appId)
				if len(appSecret) == 0 {
					return nil, errors.Forbidden("ERROR_APPID_NOT_EXISTS", "appid query parameter missing")
				}

				signature := o.signatureGetter(ctx)
				if len(signature) == 0 {
					return nil, errors.Forbidden("ERROR_MISSING_SIGNATURE", "missing query parameter: signature")
				}

				timestamp := o.timestampGetter(ctx)
				if len(signature) == 0 {
					return nil, errors.Forbidden("ERROR_MISSING_TIMESTAMP", "missing query parameter: timestamp")
				}

				u, err := time.Parse("2006-01-02 15:04:05", timestamp)
				if err != nil {
					return nil, errors.Forbidden("ERROR_INCORRECT_TIMESTAMP", "incorrect timestamp format. Should be in the form 2006-01-02 15:04:05")
				}

				t := time.Now()
				if t.Sub(u).Seconds() > float64(o.timeout) {
					return nil, errors.Forbidden("ERROR_REQUEST_TIME_TIMEOUT", "request timer timeout exceeded. Please try again")
				}

				if err := o.signatureValidator(ctx, appId, appSecret, signature, timestamp); err != nil {
					return nil, err
				}
			}
			return handler(ctx, req)
		}
	}
}
