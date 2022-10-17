package secure

import (
	"context"
)

type (
	Option func(*options)

	// options defines the config for Secure middleware.
	options struct {
		// skipper defines a function to skip middleware.
		skipper func(c context.Context) bool

		// xssProtection provides protection against cross-site scripting attack (XSS)
		// by setting the `X-XSS-Protection` header.
		// Optional. Default value "1; mode=block".
		xssProtection string `yaml:"xss_protection"`

		// contentTypeNosniff provides protection against overriding Content-Type
		// header by setting the `X-Content-Type-Options` header.
		// Optional. Default value "nosniff".
		contentTypeNosniff string `yaml:"content_type_nosniff"`

		// xFrameOptions can be used to indicate whether or not a browser should
		// be allowed to render a page in a <frame>, <iframe> or <object> .
		// Sites can use this to avoid clickjacking attacks, by ensuring that their
		// content is not embedded into other sites.provides protection against
		// clickjacking.
		// Optional. Default value "SAMEORIGIN".
		// Possible values:
		// - "SAMEORIGIN" - The page can only be displayed in a frame on the same origin as the page itself.
		// - "DENY" - The page cannot be displayed in a frame, regardless of the site attempting to do so.
		// - "ALLOW-FROM uri" - The page can only be displayed in a frame on the specified origin.
		xFrameOptions string `yaml:"x_frame_options"`

		// hstsMaxAge sets the `Strict-Transport-Security` header to indicate how
		// long (in seconds) browsers should remember that this site is only to
		// be accessed using HTTPS. This reduces your exposure to some SSL-stripping
		// man-in-the-middle (MITM) attacks.
		// Optional. Default value 0.
		hstsMaxAge int `yaml:"hsts_max_age"`

		// hstsExcludeSubdomains won't include subdomains tag in the `Strict Transport Security`
		// header, excluding all subdomains from security policy. It has no effect
		// unless HSTSMaxAge is set to a non-zero value.
		// Optional. Default value false.
		hstsExcludeSubdomains bool `yaml:"hsts_exclude_subdomains"`

		// contentSecurityPolicy sets the `Content-Security-Policy` header providing
		// security against cross-site scripting (XSS), clickjacking and other code
		// injection attacks resulting from execution of malicious content in the
		// trusted web page context.
		// Optional. Default value "".
		contentSecurityPolicy string `yaml:"content_security_policy"`

		// cspReportOnly would use the `Content-Security-Policy-Report-Only` header instead
		// of the `Content-Security-Policy` header. This allows iterative updates of the
		// content security policy by only reporting the violations that would
		// have occurred instead of blocking the resource.
		// Optional. Default value false.
		cspReportOnly bool `yaml:"csp_report_only"`

		// HSTSPreloadEnabled will add the preload tag in the `Strict Transport Security`
		// header, which enables the domain to be included in the HSTS preload list
		// maintained by Chrome (and used by Firefox and Safari): https://hstspreload.org/
		// Optional.  Default value false.
		hstsPreloadEnabled bool `yaml:"hsts_preload_enabled"`

		// ReferrerPolicy sets the `Referrer-Policy` header providing security against
		// leaking potentially sensitive request paths to third parties.
		// Optional. Default value "".
		referrerPolicy string `yaml:"referrer_policy"`
	}
)

func WithReferrerPolicy(referrerPolicy string) Option {
	return func(o *options) {
		o.referrerPolicy = referrerPolicy
	}
}

func WithSkipper(skipper func(ctx context.Context) bool) Option {
	return func(o *options) {
		o.skipper = skipper
	}
}

func WithContentSecurityPolicy(cspReportOnly bool, contentSecurityPolicy string) Option {
	return func(o *options) {
		o.cspReportOnly = cspReportOnly
		o.contentSecurityPolicy = contentSecurityPolicy
	}
}

func WithHsts(hstsPreloadEnabled bool, hstsExcludeSubdomains bool, hstsMaxAge int) Option {
	return func(o *options) {
		o.hstsPreloadEnabled = hstsPreloadEnabled
		o.hstsMaxAge = hstsMaxAge
		o.hstsExcludeSubdomains = hstsExcludeSubdomains
	}
}

func WithXFrameOptions(xFrameOptions string) Option {
	return func(o *options) {
		o.xFrameOptions = xFrameOptions
	}
}

func WithContentTypeNosniff(contentTypeNosniff string) Option {
	return func(o *options) {
		o.contentTypeNosniff = contentTypeNosniff
	}
}

func WithXSSProtection(xssProtection string) Option {
	return func(o *options) {
		o.xssProtection = xssProtection
	}
}
