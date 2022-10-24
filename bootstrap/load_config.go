package bootstrap

func LoadConfig[T any](cfg ConfigFlags) (T, func(), error) {
	c := NewConfigProvider(cfg)
	var bc T
	if err := c.Load(); err != nil {
		return bc, nil, err
	}

	if err := c.Scan(&bc); err != nil {
		return bc, nil, err
	}

	return bc, func() {
		_ = c.Close()
	}, nil
}
