package part

import (
	"net/http"
	"strings"

	"github.com/bborbe/handler_finder"
	"github.com/bborbe/handler_finder/dummy"
	"github.com/bborbe/log"
)

type handlerFinderPart struct {
	prefix  string
	handler map[string]handler_finder.HandlerFinder
}

type PartHandlerFinder interface {
	handler_finder.HandlerFinder
	RegisterHandler(part string, handler http.Handler)
	RegisterHandlerFinder(part string, handlerFinder handler_finder.HandlerFinder)
}

var logger = log.DefaultLogger

func New(prefix string) *handlerFinderPart {
	h := new(handlerFinderPart)
	h.handler = make(map[string]handler_finder.HandlerFinder)
	h.prefix = prefix
	return h
}

func (h *handlerFinderPart) RegisterHandler(part string, handler http.Handler) {
	h.RegisterHandlerFinder(part, dummy.New(handler))
}

func (h *handlerFinderPart) RegisterHandlerFinder(part string, handlerFinder handler_finder.HandlerFinder) {
	h.handler[part] = handlerFinder
}

func (h *handlerFinderPart) FindHandler(request *http.Request) http.Handler {
	hf := h.FindHandlerFinder(request)
	if hf != nil {
		return hf.FindHandler(request)
	}
	return nil
}

func (h *handlerFinderPart) FindHandlerFinder(request *http.Request) handler_finder.HandlerFinder {
	return h.FindHandlerByRequestUri(request.RequestURI)
}

func (h *handlerFinderPart) FindHandlerByRequestUri(requestUri string) handler_finder.HandlerFinder {
	rest := requestUri[len(h.prefix):]
	logger.Tracef("requestUri: %s prefix: %s => rest: %s", requestUri, h.prefix, rest)
	if len(rest) == 0 {
		return h.handler[rest]
	}
	if rest[:1] == "/" {
		pos := findEndPos(rest[1:])
		var name string
		if pos != -1 {
			name = rest[:pos+1]
		} else {
			name = rest
		}
		handler := h.handler[name]
		if handler != nil {
			return handler
		}
		return h.handler["/"]
	}
	pos := findEndPos(rest)
	var name string
	if pos != -1 {
		name = rest[:pos]
	} else {
		name = rest
	}
	return h.handler[name]
}

func findEndPos(content string) int {
	return strings.IndexFunc(content, endRunes)
}

func endRunes(r rune) bool {
	return r == '/' || r == '?'
}
