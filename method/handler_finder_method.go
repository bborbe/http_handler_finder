package method

import (
	"net/http"

	"github.com/bborbe/http_handler_finder"
	"github.com/bborbe/http_handler_finder/dummy"
)

type handlerFinderMethod struct {
	handler map[string]handler_finder.HandlerFinder
}

type MethodHandlerFinder interface {
	handler_finder.HandlerFinder
	RegisterHandler(method string, handler http.Handler)
	RegisterHandlerFinder(method string, handlerFinder handler_finder.HandlerFinder)
	GetHandlerFinder(method string) handler_finder.HandlerFinder
}

func New() *handlerFinderMethod {
	h := new(handlerFinderMethod)
	h.handler = make(map[string]handler_finder.HandlerFinder)
	return h
}

func (h *handlerFinderMethod) FindHandler(request *http.Request) http.Handler {
	hf := h.GetHandlerFinder(defaultMethod(request.Method))
	if hf != nil {
		return hf.FindHandler(request)
	}
	return nil
}

func (h *handlerFinderMethod) RegisterHandler(method string, handler http.Handler) {
	h.RegisterHandlerFinder(defaultMethod(method), dummy.New(handler))
}

func (h *handlerFinderMethod) RegisterHandlerFinder(method string, handlerFinder handler_finder.HandlerFinder) {
	h.handler[defaultMethod(method)] = handlerFinder
}

func (h *handlerFinderMethod) GetHandlerFinder(method string) handler_finder.HandlerFinder {
	return h.handler[defaultMethod(method)]
}

func defaultMethod(method string) string {
	if len(method) == 0 {
		return "GET"
	}
	return method
}
