package method

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/http_handler_finder"
)

func TestImplementsHandlerFinder(t *testing.T) {
	h := New()
	var handler *handler_finder.HandlerFinder
	err := AssertThat(h, Implements(handler).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}
