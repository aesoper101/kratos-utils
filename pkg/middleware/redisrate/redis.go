package redisrate

import (
	"context"
	"github.com/aesoper101/kratos-utils/constants"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-redis/redis_rate/v9"
	"strconv"
	"time"
)

var ErrRateLimited = errors.Forbidden("ErrRateLimited", "")

func RedisServer(limiter redis_rate.Limiter, keyFunc KeyFunc, limitFunc LimitFunc) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				res, err := limiter.Allow(ctx, keyFunc(ctx), limitFunc(ctx))
				if err != nil {
					return nil, errors.FromError(err)
				}

				h := tr.ReplyHeader()
				h.Set(constants.HeaderRateLimitRemaining, strconv.Itoa(res.Remaining))

				if res.Allowed == 0 {
					// We are rate limited.

					seconds := int(res.RetryAfter / time.Second)
					h.Set(constants.HeaderRateLimitRetryAfter, strconv.Itoa(seconds))

					// Stop processing and return the error.
					return nil, ErrRateLimited
				}
			}

			// Continue processing as normal.
			return handler(ctx, req)
		}
	}
}

type KeyFunc = func(ctx context.Context) string

type LimitFunc = func(ctx context.Context) redis_rate.Limit
