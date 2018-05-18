package mux

import (
	"net/http"

	"github.com/bborbe/http_handler_finder"
)

type handler struct {
	handlerFinder handler_finder.HandlerFinder
	errorHandler  http.Handler
}

func New(handlerFinder handler_finder.HandlerFinder, errorHandler http.Handler) *handler {
	h := new(handler)
	h.handlerFinder = handlerFinder
	h.errorHandler = errorHandler
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	handler := h.handlerFinder.FindHandler(request)
	if handler != nil {
		handler.ServeHTTP(responseWriter, request)
	} else {
		h.errorHandler.ServeHTTP(responseWriter, request)
	}
}
