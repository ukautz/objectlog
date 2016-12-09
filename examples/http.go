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
