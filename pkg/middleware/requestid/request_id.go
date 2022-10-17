package requestid

import (
	"context"
	"github.com/aesoper101/kratos-utils/constants"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func Server(opts ...Option) middleware.Middleware {
	o := &options{
		generator: DefaultGenerator,
	}

	for _, opt := range opts {
		opt(o)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				responseHeader := tr.ReplyHeader()
				requestHeader := tr.RequestHeader()

				if md, ok := metadata.FromServerContext(ctx); ok {
					requestID := md.Get(constants.HeaderXRequestID)
					if len(requestID) == 0 {
						if requestHeader.Get(constants.HeaderXRequestID) != "" {
							requestID = requestHeader.Get(constants.HeaderXRequestID)
						} else {
							requestID = o.generator(ctx)
						}
					}

					md.Set(constants.HeaderXRequestID, requestID)
					ctx = metadata.NewServerContext(ctx, md)
					responseHeader.Set(constants.HeaderXRequestID, requestID)
				}
			}

			return handler(ctx, req)
		}
	}
}
