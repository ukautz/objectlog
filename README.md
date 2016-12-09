![status](https://api.travis-ci.org/ukautz/objectlog.svg?branch=master)
[![GoDoc](https://godoc.org/gopkg.in/github.com/ukautz/objectlog?status.svg)](http://godoc.org/github.com/ukautz/objectlog)

# LogObject

**TL;DR**: Decorates Go object with logging capabilities

Writing logs is easy. Writing helpful logs is hard. This package provides a simple decoration for Go objects to make logging the right things easier.

This package supports the built-in [log package](//golang.org/pkg/log/) as well as [Logrus](//github.com/Sirupsen/logrus) out of the box.
Any other log implementation can be used as well, by writing an adapter which implements the [objectlog.ObjectLogger interface](https://godoc.org/github.com/ukautz/objectlog#ObjectLogger).

## Code pitch

```go
// bad:
// * clutters code
// * tedious to write
// * costly to change
log.Debug("[id: %s, name: %s, foo: %s]: Something happened", obj.ID(), obj.Name(), obj.Foo())

// good:
// * clean code
// * expressive
// * easy to modify
obj.LogDebug("Something happened")
```

## Install

```bash
# install standard
go get github.com/ukautz/objectlog

# install with logrus support
go get github.com/ukautz/objectlog/...
```

## Example

A simple HTTP server, which uses a wrapped request object, which is decorated by logging. Check out [more examples](https://github.com/ukautz/objectlog/tree/master/examples), if you like.

```go
package main

import (
	"fmt"
	"github.com/ukautz/objectlog"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type (

	// Requests showcases a wrapped *http.Request object with logging decoration
	Request struct {

		// decorated by ObjectLog
		*objectlog.ObjectLog

		// inherits request
		*http.Request

		// some arbitrary attribs
		uuid string
	}
)

// newUUID is dummy for function which generates an ID per request for logging
func newUUID() string {
	rand.Seed(time.Now().UnixNano())
	uuid := make([]string, 3)
	for i := 0; i < 3; i++ {
		uuid[i] = fmt.Sprintf("%03d", rand.Intn(999))
	}
	return strings.Join(uuid, "-")
}

// wrap *http.Request in local *Request object, which is decorated by logging
func newRequest(req *http.Request) *Request {
	uuid := newUUID()
	prefix := fmt.Sprintf("(uuid: %s, from: %s, path: %s, method: %s) ", uuid, req.RemoteAddr, req.URL.Path, req.Method)
	return &Request{
		ObjectLog: objectlog.NewObjectLog().SetLogPrefix(prefix),
		Request:   req,
		uuid:      uuid,
	}
}

// newHandler is helper to generate method compatible with `http.HandleFunc` while using local `*Request` object
func newHandler(cb func(rw http.ResponseWriter, req *Request)) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		req2 := newRequest(req)
		start := time.Now()
		cb(rw, req2)
		req2.LogInfo("Took %s", time.Now().Sub(start))
	}
}

// main starts HTTP server, does some demo stuff and logs duration of every received request
func main() {
	fmt.Println("Starting")
	fmt.Println("  Showcase object decoration by logging all HTTP requests")
	http.HandleFunc("/", newHandler(func(rw http.ResponseWriter, req *Request) {
		rw.Header().Add("Content-type", "text/plain")
		if req.Method == "POST" {
			req.LogWarn("POST not supported!")
			rw.WriteHeader(http.StatusNotImplemented)
			rw.Write([]byte("No POST I do"))
		} else {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("Hello there"))
		}
	}))
	http.ListenAndServe(":8000", nil)
}
```