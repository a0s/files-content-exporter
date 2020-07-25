package main

import "net/http"

type wrapper struct {
	beforeFunction  func()
	originalHandler http.Handler
}

func (ph wrapper) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ph.beforeFunction()
	ph.originalHandler.ServeHTTP(rw, r)
}

func newWrapper(beforeFunc func(), handler http.Handler) *wrapper {
	ph := wrapper{beforeFunc, handler}
	return &ph
}
