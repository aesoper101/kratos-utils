package localize

import (
	"embed"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"path/filepath"
	"text/template"
)

type (
	I18nBundle struct {
		Bundle *i18n.Bundle
	}

	Localizer struct {
		bundle    *I18nBundle
		localizer *i18n.Localizer
	}

	Config struct {
		PluralCount interface{}

		Funcs template.FuncMap
	}
)

func NewBundleFromEmbedFs(fs embed.FS, pattern string) (*I18nBundle, error) {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

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

	return &I18nBundle{Bundle: bundle}, nil
}

func NewBundleFromFile(pattern string) (*I18nBundle, error) {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

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

	return &I18nBundle{Bundle: bundle}, nil
}

func NewLocalizer(bundle *I18nBundle, langs ...string) *Localizer {
	localizer := i18n.NewLocalizer(bundle.Bundle, langs...)

	return &Localizer{
		localizer: localizer,
	}
}

func (l *Localizer) T(messageId string, data interface{}, args ...Config) string {
	cfg := &i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: data,
	}

	if len(args) > 0 {
		cfg.PluralCount = args[0].PluralCount
		cfg.Funcs = args[0].Funcs
	}

	msg, err := l.localizer.Localize(cfg)
	if err != nil {
		return ""
	}

	return msg
}
