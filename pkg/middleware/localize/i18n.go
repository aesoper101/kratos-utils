package localize

import (
	"embed"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/aesoper101/go-utils/templatex"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"sync"
	"text/template"
)

type (
	I18nBundle struct {
		bundle *i18n.Bundle
		reMux  sync.RWMutex
		funcs  template.FuncMap
	}

	Localizer struct {
		bundle    *I18nBundle
		localizer *i18n.Localizer
		funcs     template.FuncMap
	}
)

func NewBundleFromEmbedFs(fs embed.FS, pattern string) (*I18nBundle, error) {
	bundle := i18n.NewBundle(language.Chinese)
	setDefaultUnmarshalFunc(bundle)

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, m := range matches {
		_, err := bundle.LoadMessageFileFS(fs, m)
		if err != nil {
			return nil, err
		}
	}

	return &I18nBundle{
		bundle: bundle,
		funcs:  templatex.DefaultTemplateFuncMap,
	}, nil
}

func NewBundleFromFile(pattern string) (*I18nBundle, error) {
	bundle := i18n.NewBundle(language.Chinese)
	setDefaultUnmarshalFunc(bundle)

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, m := range matches {
		_, err := bundle.LoadMessageFile(m)
		if err != nil {
			return nil, err
		}
	}

	return &I18nBundle{
		bundle: bundle,
		funcs:  templatex.DefaultTemplateFuncMap,
	}, nil
}

func (b *I18nBundle) RegisterUnmarshalFunc(format string, unmarshalFunc i18n.UnmarshalFunc) {
	b.bundle.RegisterUnmarshalFunc(format, unmarshalFunc)
}

func (b *I18nBundle) RegisterFuncs(funcs template.FuncMap) {
	b.reMux.Lock()
	defer b.reMux.Unlock()

	for k, v := range funcs {
		b.funcs[k] = v
	}
}

func setDefaultUnmarshalFunc(bundle *i18n.Bundle) {
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
}

func NewLocalizer(bundle *I18nBundle, langs ...string) *Localizer {
	localizer := i18n.NewLocalizer(bundle.bundle, langs...)

	return &Localizer{
		localizer: localizer,
		funcs:     bundle.funcs,
	}
}

func (l *Localizer) T(messageId string, data interface{}, pluralCounts ...interface{}) string {
	cfg := &i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: data,
		Funcs:        l.funcs,
	}

	if len(pluralCounts) > 0 {
		cfg.PluralCount = pluralCounts[0]
	}

	msg, err := l.localizer.Localize(cfg)
	if err != nil {
		return ""
	}

	return msg
}
