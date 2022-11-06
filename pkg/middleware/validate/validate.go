package validate

import (
	"context"
	"github.com/aesoper101/kratos-utils/pkg/middleware/translator"
	"github.com/aesoper101/kratos-utils/protobuf/types/httppb"
	protoValidator "github.com/aesoper101/protoc-gen-govalidate/validator"
	"github.com/go-kratos/kratos/v2/middleware"
	ut "github.com/go-playground/universal-translator"
	govalidator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

type validator interface {
	Validate() error
	ValidateCtx(ctx context.Context) error
}

type (
	Option func(*options)

	options struct {
		registerDefaultTranslations func(validate *govalidator.Validate, trans ut.Translator)
	}
)

func WithRegisterDefaultTranslations(fn func(validate *govalidator.Validate, trans ut.Translator)) Option {
	return func(o *options) {
		o.registerDefaultTranslations = fn
	}
}

// Validator is a validator middleware.
func Validator(opts ...Option) middleware.Middleware {
	o := &options{
		registerDefaultTranslations: registerDefaultTranslations,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			trans := translator.FromTranslatorContext(ctx)
			unTrans := trans.RawTranslator()
			o.registerDefaultTranslations(protoValidator.Validator(), unTrans)

			if v, ok := req.(validator); ok {
				if err := v.Validate(); err != nil {

					errs, ok := err.(govalidator.ValidationErrors)
					if ok && len(errs) > 0 {
						metadata := map[string]string{}
						for _, err := range errs {
							metadata[err.Field()] = err.Translate(unTrans)
						}

						return nil, httppb.ErrorBadRequest(errs[0].Translate(unTrans)).WithCause(err).WithMetadata(metadata)
					}

					return nil, httppb.ErrorBadRequest(err.Error()).WithCause(err)
				}
			}
			return handler(ctx, req)
		}
	}
}

func registerDefaultTranslations(validate *govalidator.Validate, trans ut.Translator) {
	switch trans.Locale() {
	case "en":
		_ = enTranslations.RegisterDefaultTranslations(validate, trans)
		return
	case "zh":
		_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
		return
	default:
		_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
	}
}
