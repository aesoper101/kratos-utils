package translator

import (
	"context"
	"embed"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/tdewolff/parse/v2/buffer"
	fs2 "io/fs"
	"log"
)

func NewTranslatorFromFs(fs embed.FS, localePath string, supportedLocales ...locales.Translator) (*ut.UniversalTranslator, error) {
	trans := NewUniversalTranslator(supportedLocales...)

	err := fs2.WalkDir(fs, localePath, func(path string, d fs2.DirEntry, err error) error {
		if !d.IsDir() {
			file, err := fs2.ReadFile(fs, path)
			if err != nil {
				return err
			}

			err = trans.ImportByReader(ut.FormatJSON, buffer.NewReader(file))
			if err != nil {

				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if err := trans.VerifyTranslations(); err != nil {
		return nil, err
	}

	return trans, nil
}

func NewUniversalTranslator(supportedLocales ...locales.Translator) *ut.UniversalTranslator {
	supportedLocales = append(supportedLocales, zh.New(), en.New())

	return ut.New(zh.New(), supportedLocales...)
}

type transKey struct{}

func newTranslatorContext(ctx context.Context, trans *ut.UniversalTranslator, currencyFn func(locale string) currency.Type, locale string) context.Context {
	newTrans := newTranslator(trans, locale)

	return context.WithValue(ctx, transKey{}, &translator{
		Translator: newTrans.(locales.Translator),
		trans:      newTrans,
		currencyFn: currencyFn,
	})
}

func FromTranslatorContext(ctx context.Context) Translator {
	return ctx.Value(transKey{}).(Translator)
}

func newTranslator(trans *ut.UniversalTranslator, locale string) ut.Translator {
	translator, found := trans.GetTranslator(locale)
	if found {
		return translator
	}

	findTranslator, _ := trans.FindTranslator(locale)
	return findTranslator
}

type Translator interface {
	locales.Translator

	T(key interface{}, params ...string) string

	C(key interface{}, num float64, digits uint64, param string) string

	O(key interface{}, num float64, digits uint64, param string) string

	R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string

	// Currency returns the type used by the given locale.
	Currency() currency.Type

	RawTranslator() ut.Translator
}

type translator struct {
	locales.Translator
	trans      ut.Translator
	currencyFn func(locale string) currency.Type
}

func (t *translator) T(key interface{}, params ...string) string {

	s, err := t.trans.T(key, params...)
	if err != nil {
		log.Printf("issue translating key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) C(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.Printf("issue translating cardinal key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) O(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.Printf("issue translating ordinal key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string {

	s, err := t.trans.R(key, num1, digits1, num2, digits2, param1, param2)
	if err != nil {
		log.Printf("issue translating range key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *translator) Currency() currency.Type {
	if t.currencyFn != nil {
		currencyType := t.currencyFn(t.Locale())
		if currencyType >= 0 {
			return currencyType
		}
	}
	switch t.Locale() {
	case "en":
		return currency.USD
	case "zh":
		return currency.CNY
	default:
		return currency.CNY
	}
}

func (t *translator) RawTranslator() ut.Translator {
	return t.trans
}

var _ Translator = (*translator)(nil)

type (
	Option func(*options)

	options struct {
		trans      *ut.UniversalTranslator
		localeFn   func(ctx context.Context) string
		currencyFn func(locale string) currency.Type
	}
)

func WithUniversalTranslator(trans *ut.UniversalTranslator) Option {
	return func(o *options) {
		o.trans = trans
	}
}

func WithLocaleFn(fn func(ctx context.Context) string) Option {
	return func(o *options) {
		o.localeFn = fn
	}
}

func WithCurrencyFn(fn func(locale string) currency.Type) Option {
	return func(o *options) {
		o.currencyFn = fn
	}
}

func Translate(opts ...Option) middleware.Middleware {
	o := &options{
		trans: NewUniversalTranslator(),
		localeFn: func(ctx context.Context) string {
			if tr, ok := transport.FromServerContext(ctx); ok {
				return tr.RequestHeader().Get("Accept-Language")
			}
			return "zh"
		},
	}

	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			locale := o.localeFn(ctx)
			ctx = newTranslatorContext(ctx, o.trans, o.currencyFn, locale)
			return handler(ctx, req)
		}
	}
}
