package secure

import (
	"context"
	"fmt"
	"github.com/aesoper101/kratos-utils/constants"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// Server only use in http.Middleware
func Server(opts ...Option) middleware.Middleware {
	config := &options{
		skipper: func(c context.Context) bool {
			return false
		},
		xssProtection:      "1; mode=block",
		contentTypeNosniff: "nosniff",
		xFrameOptions:      "SAMEORIGIN",
		hstsPreloadEnabled: false,
	}

	for _, opt := range opts {
		opt(config)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if config.skipper(ctx) {
				return handler(ctx, req)
			}
			if tr, ok := transport.FromServerContext(ctx); ok {
				reqHeader := tr.RequestHeader()
				respHeader := tr.ReplyHeader()

				if config.xssProtection != "" {
					respHeader.Set(constants.HeaderXXSSProtection, config.xssProtection)
				}
				if config.contentTypeNosniff != "" {
					respHeader.Set(constants.HeaderXContentTypeOptions, config.contentTypeNosniff)
				}
				if config.xFrameOptions != "" {
					respHeader.Set(constants.HeaderXFrameOptions, config.xFrameOptions)
				}

				if (reqHeader.Get(constants.HeaderXForwardedProto) == "https") && config.hstsMaxAge != 0 {
					subdomains := ""
					if !config.hstsExcludeSubdomains {
						subdomains = "; includeSubdomains"
					}
					if config.hstsPreloadEnabled {
						subdomains = fmt.Sprintf("%s; preload", subdomains)
					}
					respHeader().Set(constants.HeaderStrictTransportSecurity, fmt.Sprintf("max-age=%d%s", config.hstsMaxAge, subdomains))
				}

				if config.contentSecurityPolicy != "" {
					if config.cspReportOnly {
						respHeader().Set(constants.HeaderContentSecurityPolicyReportOnly, config.contentSecurityPolicy)
					} else {
						respHeader().Set(constants.HeaderContentSecurityPolicy, config.contentSecurityPolicy)
					}
				}
				if config.referrerPolicy != "" {
					respHeader().Set(constants.HeaderReferrerPolicy, config.referrerPolicy)
				}
			}

			return handler(ctx, req)
		}
	}
}
