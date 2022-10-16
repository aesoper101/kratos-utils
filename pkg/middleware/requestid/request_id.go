package requestid

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	uuid "github.com/satori/go.uuid"
)

const HeaderRequestId = "X-Request-Id"

func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				responseHeader := tr.ReplyHeader()
				requestHeader := tr.RequestHeader()

				if md, ok := metadata.FromServerContext(ctx); ok {
					requestID := md.Get(HeaderRequestId)
					if len(requestID) == 0 {
						if requestHeader.Get(HeaderRequestId) != "" {
							requestID = requestHeader.Get(HeaderRequestId)
						} else {
							requestID = uuid.NewV4().String()
						}
					}

					md.Set(HeaderRequestId, requestID)
					ctx = metadata.NewServerContext(ctx, md)
					responseHeader.Set(HeaderRequestId, requestID)
				}
			}

			return handler(ctx, req)
		}
	}
}
