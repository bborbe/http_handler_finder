package static

import "net/http"

type handler struct {
	content    string
	returnCode int
}

func New(content string) *handler {
	return NewWithReturnCode(content, http.StatusOK)
}

func NewWithReturnCode(content string, returnCode int) *handler {
	h := new(handler)
	h.content = content
	h.returnCode = returnCode
	return h
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(h.returnCode)
	responseWriter.Write([]byte(h.content))
}
