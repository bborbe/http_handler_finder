package path

import "net/http"

type handlerFinderPath map[string]http.Handler

func New() *handlerFinderPath {
	h := make(handlerFinderPath)
	return &h
}

func (h handlerFinderPath) RegisterHandler(path string, handler http.Handler) {
	h[path] = handler
}

func (h handlerFinderPath) FindHandler(request *http.Request) http.Handler {
	return h[request.RequestURI]
}
