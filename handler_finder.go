package handler_finder

import "net/http"

type HandlerFinder interface {
	FindHandler(request *http.Request) http.Handler
}
