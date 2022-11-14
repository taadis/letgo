package cache

import (
	"time"
)

// Options represents the options for the cache.
type Options struct {
	Expiration time.Duration
	Items      map[string]Item
	Address    string
}

type Option func(*Options)

// NewOptions returns a new options struct.
func NewOptions(opts ...Option) Options {
	options := Options{
		Expiration: DefaultExpiration,
		Items:      make(map[string]Item),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func WithExpiration(d time.Duration) Option {
	return func(o *Options) {
		o.Expiration = d
	}
}

// WithAddress sets the cache service address or connection information.
func WithAddress(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

func WithItems(i map[string]Item) Option {
	return func(o *Options) {
		o.Items = i
	}
}
