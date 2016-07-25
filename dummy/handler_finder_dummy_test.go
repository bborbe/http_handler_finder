package dummy

import (
	"testing"

	"net/http"

	. "github.com/bborbe/assert"
	"github.com/bborbe/handler_finder"
	"github.com/bborbe/server/handler/static"
)

func TestImplementsHandlerFinder(t *testing.T) {
	hf := New(static.NewHandlerStaticContent("test"))
	var handlerFinder *handler_finder.HandlerFinder
	err := AssertThat(hf, Implements(handlerFinder).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindHandlerNotNil(t *testing.T) {
	hf := New(static.NewHandlerStaticContent("test"))
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
