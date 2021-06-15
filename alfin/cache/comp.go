package cache

import "context"

type cacheCtxKey struct{}

// FromContext returns a Cache stored inside a context, or nil if there isn't one.
func FromContext(ctx context.Context) Cache {
	c, _ := ctx.Value(cacheCtxKey{}).(Cache)
	return c
}

// NewContext returns a new context with the given Client attached.
func NewContext(parent context.Context) context.Context {
	return context.WithValue(parent, cacheCtxKey{},NewLRU(10000))
}

