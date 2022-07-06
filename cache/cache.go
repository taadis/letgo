package cache

import (
	"context"
	"crypto/tls"
	"errors"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key not found in cache")
	ErrExpired     = errors.New("has expired")
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Del(ctx context.Context, key string) error
}

type Options struct {
	Timeout time.Duration
	TLS     *tls.Config
	Version string
	Address string
	// more...
}

type Option func(*Options)

func WithTimeout(d time.Duration) Option {
	return func(o *Options) {
		o.Timeout = d
	}
}

func WithTLS(c *tls.Config) Option {
	return func(o *Options) {
		o.TLS = c
	}
}

func WithVersion(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

func WithAddress(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}
