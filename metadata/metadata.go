// Package metadata is a way of defining message headers.
package metadata

import "context"

type metaKey struct{}

type Metadata map[string]string

func Copy(md Metadata) Metadata {
	mdata := make(Metadata)
	for k, v := range md {
		mdata[k] = v
	}
	return mdata
}

func FromContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(metaKey{}).(Metadata)
	return md, ok
}

func NewContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, metaKey{}, md)
}

func WithMetadata(ctx context.Context, md Metadata) context.Context {
	cmd, ok := ctx.Value(metaKey{}).(Metadata)
	if ok {
		for k, v := range cmd {
			if _, ok := md[k]; !ok {
				md[k] = v
			}
		}
	}
	return context.WithValue(ctx, metaKey{}, md)
}
