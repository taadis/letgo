package registry

import (
	"crypto/tls"
	"time"
)

type Options struct {
	Address   []string
	Timeout   time.Duration
	Secure    bool
	TLSConfig *tls.Config
}

type Option func(*Options)

type RegistryOptions struct {
	TTL time.Duration
}

type RegistryOption func(*RegistryOptions)

type WatchOptions struct {
	Service string
}

type WatchOption func(*WatchOptions)

type DeregisterOptions struct {
}

type GetOptions struct {
}

type ListOptions struct {
}

func WithAddress(address ...string) Option {
	return func(o *Options) {
		o.Address = address
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func WithSecure(b bool) Option {
	return func(o *Options) {
		o.Secure = b
	}
}

func WithTTL(d time.Duration) RegistryOption {
	return func(o *RegistryOptions) {
		o.TTL = d
	}
}

func WithService(name string) WatchOption {
	return func(o *WatchOptions) {
		o.Service = name
	}
}
