package rest

import (
	"net/http"

	handler_finder_method "github.com/bborbe/http_handler_finder/method"
	handler_finder_part "github.com/bborbe/http_handler_finder/part"
)

type RestHandlerFinder interface {
	RegisterListHandler(handler http.Handler)
	RegisterGetHandler(handler http.Handler)
	RegisterCreateHandler(handler http.Handler)
	RegisterDeleteHandler(handler http.Handler)
	RegisterUpdateHandler(handler http.Handler)
	RegisterPatchHandler(handler http.Handler)
	RegisterHandler(method string, part string, handler http.Handler)
}

type restHandlerFinder struct {
	prefix              string
	partHandlerFinder   map[string]handler_finder_part.PartHandlerFinder
	methodHandlerFinder handler_finder_method.MethodHandlerFinder
}

func New(prefix string) *restHandlerFinder {
	h := new(restHandlerFinder)
	h.prefix = prefix
	h.methodHandlerFinder = handler_finder_method.New()
	h.partHandlerFinder = make(map[string]handler_finder_part.PartHandlerFinder)
	return h
}

func (h *restHandlerFinder) FindHandler(request *http.Request) http.Handler {
	return h.methodHandlerFinder.FindHandler(request)
}

func (h *restHandlerFinder) getPartHandlerFinderByMethod(method string) handler_finder_part.PartHandlerFinder {
	handlerFinder := h.partHandlerFinder[method]
	if handlerFinder == nil {
		handlerFinder = handler_finder_part.New(h.prefix)
		h.partHandlerFinder[method] = handlerFinder
		h.methodHandlerFinder.RegisterHandlerFinder(method, handlerFinder)
	}
	return handlerFinder
}

func (h *restHandlerFinder) RegisterHandler(method string, part string, handler http.Handler) {
	h.getPartHandlerFinderByMethod(method).RegisterHandler(part, handler)
}

func (h *restHandlerFinder) RegisterListHandler(handler http.Handler) {
	h.RegisterHandler("GET", "", handler)
}

func (h *restHandlerFinder) RegisterGetHandler(handler http.Handler) {
	h.RegisterHandler("GET", "/", handler)
}

func (h *restHandlerFinder) RegisterCreateHandler(handler http.Handler) {
	h.RegisterHandler("POST", "", handler)
	h.RegisterHandler("POST", "/", handler)
}

func (h *restHandlerFinder) RegisterDeleteHandler(handler http.Handler) {
	h.RegisterHandler("DELETE", "", handler)
	h.RegisterHandler("DELETE", "/", handler)
}

func (h *restHandlerFinder) RegisterUpdateHandler(handler http.Handler) {
	h.RegisterHandler("PUT", "", handler)
	h.RegisterHandler("PUT", "/", handler)
}

func (h *restHandlerFinder) RegisterPatchHandler(handler http.Handler) {
	h.RegisterHandler("PATCH", "", handler)
	h.RegisterHandler("PATCH", "/", handler)
}
