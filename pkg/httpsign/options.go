package httpsign

import "context"

type Option func(*options)

type AppSecretGetter func(appId string) string

type AppIdGetter func(ctx context.Context) string

type SignatureGetter func(ctx context.Context) string

type TimestampGetter func(ctx context.Context) string

type SignatureValidator func(ctx context.Context, appId, secret, signature, timestamp string) error

type options struct {
	secretGetter       AppSecretGetter
	appIdGetter        AppIdGetter
	signatureGetter    SignatureGetter
	timestampGetter    TimestampGetter
	timeout            int
	signatureValidator SignatureValidator
}

func WithAppSecretGetter(getter AppSecretGetter) Option {
	return func(o *options) {
		o.secretGetter = getter
	}
}

func WithAppIdGetter(getter AppIdGetter) Option {
	return func(o *options) {
		o.appIdGetter = getter
	}
}

func WithSecretGetter(getter SignatureGetter) Option {
	return func(o *options) {
		o.signatureGetter = getter
	}
}

func WithTimestampGetter(getter TimestampGetter) Option {
	return func(o *options) {
		o.timestampGetter = getter
	}
}

func WithSignatureValidator(validator SignatureValidator) Option {
	return func(o *options) {
		o.signatureValidator = validator
	}
}
