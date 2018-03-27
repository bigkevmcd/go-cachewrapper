package cachewrapper

import (
	"net/http"
	"time"
)

type cacheWrapper struct {
	opts CacheOptions
}

func CacheWrapper() *cacheWrapper {
	return &cacheWrapper{opts: CacheOptions{}}
}

// Immutable configures the max-age pragma to be one year in the future.
func (cw *cacheWrapper) Immutable() *cacheWrapper {
	cw.opts.Immutable = true
	return cw
}

// MaxAge configures the maximum-age cache option.
func (cw *cacheWrapper) MaxAge(d time.Duration) *cacheWrapper {
	cw.opts.MaxAge = d
	return cw
}

// NoTransform configures the the no-transform pragma.
func (cw *cacheWrapper) NoTransform() *cacheWrapper {
	cw.opts.NoTransform = true
	return cw
}

// Private configures the private cache option.
func (cw *cacheWrapper) Private() *cacheWrapper {
	cw.opts.Private = true
	return cw
}

// NoCache configures the no-cache pragma.
func (cw *cacheWrapper) NoCache() *cacheWrapper {
	cw.opts.NoCache = true
	return cw
}

// NoStore configures the no-store pragma.
func (cw *cacheWrapper) NoStore() *cacheWrapper {
	cw.opts.NoStore = true
	return cw
}

// MustRevalidate configures the must-revalidate pragma.
func (cw *cacheWrapper) MustRevalidate() *cacheWrapper {
	cw.opts.MustRevalidate = true
	return cw
}

// ProxyRevalidate configures the proxy-revalidate pragma.
func (cw *cacheWrapper) ProxyRevalidate() *cacheWrapper {
	cw.opts.ProxyRevalidate = true
	return cw
}

// SharedMaxAge configures the s-maxage pragma.
func (cw *cacheWrapper) SharedMaxAge(d time.Duration) *cacheWrapper {
	cw.opts.SharedMaxAge = d
	return cw
}

// Wrap takes an http.Handler and wraps it so it will have cache-control headers
// added.
func (cw *cacheWrapper) Wrap(f http.Handler) *CacheControl {
	return Cached(f, Config(cw.opts))
}
