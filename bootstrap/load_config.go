package bootstrap

import "github.com/go-kratos/kratos/v2/config"

type ConfigLoader[T any] struct {
	loader config.Config
	cfg    *T
}

func NewConfigLoader[T any](cfg ConfigFlags) (*ConfigLoader[T], func(), error) {
	c := NewConfigProvider(cfg)
	loader := &ConfigLoader[T]{loader: c, cfg: new(T)}

	err := loader.loadConfig()
	if err != nil {
		return nil, func() {}, err
	}

	return loader, func() {
		if c != nil {
			_ = c.Close()
		}
	}, nil
}

func (c *ConfigLoader[T]) Reload() error {
	return c.loadConfig()
}

func (c *ConfigLoader[T]) GetConfig() *T {
	return c.cfg
}

func (c *ConfigLoader[T]) loadConfig() error {
	if err := c.loader.Load(); err != nil {
		return err
	}

	if err := c.loader.Scan(&c.cfg); err != nil {
		return err
	}

	return nil
}

func (c *ConfigLoader[T]) AddWatch(key string, observer config.Observer) error {
	return c.loader.Watch(key, observer)
}
