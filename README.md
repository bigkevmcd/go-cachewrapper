go-cachewrapper
===============

Trivial wrapper for http.Handler that applies Cache-Control headers in the style of DropWizard.

Example
-------

```Go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bigkevmcd/go-cachewrapper"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

const cacheTime = time.Minute * 60

func main() {
	http.Handle("/", cachewrapper.Cached(http.HandlerFunc(helloWorld), cachewrapper.CacheOptions{MaxAge: cacheTime, NoTransform: true}))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

```
$ curl -i http://localhost:8000/
HTTP/1.1 200 OK
Cache-Control: no-transform, max-age=3600
Date: Tue, 29 Apr 2014 12:14:09 GMT
Content-Length: 14
Content-Type: text/plain; charset=utf-8

Hello, World!
```
