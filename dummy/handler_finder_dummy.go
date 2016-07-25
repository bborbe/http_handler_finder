package dummy

import (
	"net/http"
)

type handlerFinderDummy struct {
	handler http.Handler
}

func New(handler http.Handler) *handlerFinderDummy {
	h := new(handlerFinderDummy)
	h.handler = handler
	return h
}

func (h *handlerFinderDummy) FindHandler(request *http.Request) http.Handler {
	return h.handler
}
