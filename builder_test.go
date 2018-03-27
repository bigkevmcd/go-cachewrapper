package cachewrapper

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {}

func assertPragmas(t *testing.T, cc *CacheControl, wanted string) {
	r, _ := http.NewRequest("POST", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	cc.ServeHTTP(w, r)

	if pragmas := w.Header().Get("Cache-Control"); pragmas != wanted {
		t.Fatalf("Cache-Control header: got %s, wanted '%s'", pragmas, wanted)
	}

}

func TestBuildWithMultipleOptions(t *testing.T) {
	w := CacheWrapper().Immutable().Private().Wrap(http.HandlerFunc(handler))

	wanted := "private, max-age=31536000"
	assertPragmas(t, w, wanted)
}

func TestBuildImmutable(t *testing.T) {
	w := CacheWrapper().Immutable().Wrap(http.HandlerFunc(handler))

	wanted := "max-age=31536000"
	assertPragmas(t, w, wanted)
}

func TestBuildMaxAge(t *testing.T) {
	w := CacheWrapper().MaxAge(time.Second * 60).Wrap(http.HandlerFunc(handler))

	wanted := "max-age=60"
	assertPragmas(t, w, wanted)
}

func TestBuildPrivate(t *testing.T) {
	w := CacheWrapper().Private().Wrap(http.HandlerFunc(handler))

	wanted := "private"
	assertPragmas(t, w, wanted)
}

func TestBuildNoCache(t *testing.T) {
	w := CacheWrapper().NoCache().Wrap(http.HandlerFunc(handler))

	wanted := "no-cache"
	assertPragmas(t, w, wanted)
}

func TestBuildNoStore(t *testing.T) {
	w := CacheWrapper().NoStore().Wrap(http.HandlerFunc(handler))

	wanted := "no-store"
	assertPragmas(t, w, wanted)
}

func TestBuildMustRevalidate(t *testing.T) {
	w := CacheWrapper().MustRevalidate().Wrap(http.HandlerFunc(handler))

	wanted := "must-revalidate"
	assertPragmas(t, w, wanted)
}

func TestBuildProxyRevalidate(t *testing.T) {
	w := CacheWrapper().ProxyRevalidate().Wrap(http.HandlerFunc(handler))

	wanted := "proxy-revalidate"
	assertPragmas(t, w, wanted)
}

func TestBuildSharedMaxAge(t *testing.T) {
	w := CacheWrapper().SharedMaxAge(time.Minute * 180).Wrap(http.HandlerFunc(handler))

	wanted := "s-maxage=10800"
	assertPragmas(t, w, wanted)
}
