package main

import "net/http"

type HandlerWrapper struct {
	beforeFunction  func()
	originalHandler http.Handler
}

func (ph HandlerWrapper) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ph.beforeFunction()
	ph.originalHandler.ServeHTTP(rw, r)
}

func NewHandlerWrapper(beforeFunc func(), handler http.Handler) *HandlerWrapper {
	ph := HandlerWrapper{beforeFunc, handler}
	return &ph
}
