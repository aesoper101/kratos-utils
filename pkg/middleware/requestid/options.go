package requestid

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type Generator func(ctx context.Context) string

type Option func(*options)

type options struct {
	generator Generator
}

func WithGenerator(generator Generator) Option {
	return func(o *options) {
		o.generator = generator
	}
}

func DefaultGenerator(_ context.Context) string {
	return uuid.NewV4().String()
}
