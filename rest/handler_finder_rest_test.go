package rest

import (
	"testing"

	"net/http"

	. "github.com/bborbe/assert"
	"github.com/bborbe/http/mock"
	"github.com/bborbe/http_handler/static"
	"github.com/bborbe/http_handler_finder"
)

func TestImplementsRestHandlerFinder(t *testing.T) {
	r := New("/test")
	var i *RestHandlerFinder
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsHandlerFinder(t *testing.T) {
	hf := New("/test")
	hf.RegisterCreateHandler(static.New("create"))
	hf.RegisterGetHandler(static.New("get"))
	hf.RegisterDeleteHandler(static.New("delete"))
	hf.RegisterUpdateHandler(static.New("update"))
	hf.RegisterListHandler(static.New("list"))
	hf.RegisterPatchHandler(static.New("patch"))
	var handler *handler_finder.HandlerFinder
	err := AssertThat(hf, Implements(handler).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	hf := New("/test")
	hf.RegisterCreateHandler(static.New("create"))
	hf.RegisterGetHandler(static.New("get"))
	hf.RegisterDeleteHandler(static.New("delete"))
	hf.RegisterUpdateHandler(static.New("update"))
	hf.RegisterListHandler(static.New("list"))
	hf.RegisterPatchHandler(static.New("patch"))
	r, _ := http.NewRequest("GET", "http://www.example.com/test/123", nil)
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(resp.String(), Is("get"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestList(t *testing.T) {
	hf := New("/test")
	hf.RegisterCreateHandler(static.New("create"))
	hf.RegisterGetHandler(static.New("get"))
	hf.RegisterDeleteHandler(static.New("delete"))
	hf.RegisterUpdateHandler(static.New("update"))
	hf.RegisterListHandler(static.New("list"))
	hf.RegisterPatchHandler(static.New("patch"))
	r, _ := http.NewRequest("GET", "http://www.example.com/test", nil)
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(resp.String(), Is("list"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestListNoMethod(t *testing.T) {
	hf := New("/test")
	hf.RegisterCreateHandler(static.New("create"))
	hf.RegisterGetHandler(static.New("get"))
	hf.RegisterDeleteHandler(static.New("delete"))
	hf.RegisterUpdateHandler(static.New("update"))
	hf.RegisterListHandler(static.New("list"))
	hf.RegisterPatchHandler(static.New("patch"))
	r, _ := http.NewRequest("", "http://www.example.com/test", nil)
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(resp.String(), Is("list"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	hf := New("/test")
	hf.RegisterCreateHandler(static.New("create"))
	hf.RegisterGetHandler(static.New("get"))
	hf.RegisterDeleteHandler(static.New("delete"))
	hf.RegisterUpdateHandler(static.New("update"))
	hf.RegisterListHandler(static.New("list"))
	hf.RegisterPatchHandler(static.New("patch"))
	r, _ := http.NewRequest("POST", "http://www.example.com/test", nil)
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(resp.String(), Is("create"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	hf := New("/test")
	hf.RegisterCreateHandler(static.New("create"))
	hf.RegisterGetHandler(static.New("get"))
	hf.RegisterDeleteHandler(static.New("delete"))
	hf.RegisterUpdateHandler(static.New("update"))
	hf.RegisterListHandler(static.New("list"))
	hf.RegisterPatchHandler(static.New("patch"))
	r, _ := http.NewRequest("PUT", "http://www.example.com/test/123", nil)
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(resp.String(), Is("update"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPatch(t *testing.T) {
	hf := New("/test")
	hf.RegisterCreateHandler(static.New("create"))
	hf.RegisterGetHandler(static.New("get"))
	hf.RegisterDeleteHandler(static.New("delete"))
	hf.RegisterUpdateHandler(static.New("update"))
	hf.RegisterListHandler(static.New("list"))
	hf.RegisterPatchHandler(static.New("patch"))
	r, _ := http.NewRequest("PATCH", "http://www.example.com/test/123", nil)
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(resp.String(), Is("patch"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	hf := New("/test")
	hf.RegisterCreateHandler(static.New("create"))
	hf.RegisterGetHandler(static.New("get"))
	hf.RegisterDeleteHandler(static.New("delete"))
	hf.RegisterUpdateHandler(static.New("update"))
	hf.RegisterListHandler(static.New("list"))
	hf.RegisterPatchHandler(static.New("patch"))
	r, _ := http.NewRequest("DELETE", "http://www.example.com/test/123", nil)
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(resp.String(), Is("delete"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterCustomerHandler(t *testing.T) {
	hf := New("/test")
	hf.RegisterHandler("POST", "/verify", static.New("verify"))
	r, _ := http.NewRequest("POST", "http://www.example.com/test/verify", nil)
	h := hf.FindHandler(r)
	err := AssertThat(h, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	resp := mock.NewHttpResponseWriterMock()
	h.ServeHTTP(resp, r)
	err = AssertThat(resp.String(), Is("verify"))
	if err != nil {
		t.Fatal(err)
	}
}
