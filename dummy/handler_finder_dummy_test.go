package dummy

import (
	"testing"

	"net/http"

	. "github.com/bborbe/assert"
	"github.com/bborbe/http_handler/static"
	"github.com/bborbe/http_handler_finder"
)

func TestImplementsHandlerFinder(t *testing.T) {
	hf := New(static.New("test"))
	var handlerFinder *handler_finder.HandlerFinder
	err := AssertThat(hf, Implements(handlerFinder).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindHandlerNotNil(t *testing.T) {
	hf := New(static.New("test"))
	h := hf.FindHandler(&http.Request{})
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindHandlerNil(t *testing.T) {
	hf := New(nil)
	h := hf.FindHandler(&http.Request{})
	err := AssertThat(h, NilValue())
	if err != nil {
		t.Fatal(err)
	}
}
