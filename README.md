go-cachewrapper
===============

Trivial wrapper for http.Handler that applies Cache-Control headers in the style of DropWizard.

This uses a range of option functions to configure the cache header pragmas.

Example
-------

```Go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	cw "github.com/bigkevmcd/go-cachewrapper"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

const cacheTime = time.Minute * 60

func main() {
	http.Handle("/", cw.Cached(http.HandlerFunc(helloWorld), cw.MaxAge(cacheTime), cw.NoTransform())
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

It's easy to see the result using curl.

```
$ curl -i http://localhost:8000/
HTTP/1.1 200 OK
Cache-Control: no-transform, max-age=3600
Date: Tue, 29 Apr 2014 12:14:09 GMT
Content-Length: 14
Content-Type: text/plain; charset=utf-8

Hello, World!
```

Backwards compatibility
-----------------------

If you need backwards compatibility with `CacheOptions` configuration, you can
wrap it in a `Config` option with...

```go
	http.Handle("/", cw.Cached(http.HandlerFunc(helloWorld), cw.{Config(cw.CacheOptions{MaxAge: time.Hour * 24 * 13, NoTransform: true})))
}
```
