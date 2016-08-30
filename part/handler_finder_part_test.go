package part

import (
	"testing"

	. "github.com/bborbe/assert"
	server_mock "github.com/bborbe/http/mock"
	"github.com/bborbe/http_handler/static"
	"github.com/bborbe/http_handler_finder"
)

func TestImplementsHandlerFinder(t *testing.T) {
	hf := New("/test")
	var handlerFinder *handler_finder.HandlerFinder
	err := AssertThat(hf, Implements(handlerFinder).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRoot(t *testing.T) {
	hf := New("")
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NilValue().Message("no handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	hf.RegisterHandler("/api", static.New("/api"))
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/foo")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NilValue().Message("no handler for /foo registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/v1")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSub(t *testing.T) {
	hf := New("/api")
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/test")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NilValue().Message("no handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	hf.RegisterHandler("/test", static.New("/test"))
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/test")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/foo")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NilValue().Message("no handler for /foo registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/test/foo")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestEmpty(t *testing.T) {
	hf := New("/api/v1/task")
	hf.RegisterHandler("", static.New(""))
	hf.RegisterHandler("/", static.New("/"))
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/v1/task")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/v1/task/")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/v1/task/123")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSubWithoutSlash(t *testing.T) {
	hf := New("/hello")
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/helloworld")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NilValue().Message("no handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	hf.RegisterHandler("world", static.New("world"))
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/helloworld")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/helloworld/foo")
		if err != nil {
			t.Error(err)
		}
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestWithQuestionMark(t *testing.T) {
	content := "TestHandlerContent"
	hf := New("/api")
	hf.RegisterHandler("/test", static.New(content))
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/test")
		if err != nil {
			t.Error(err)
		}
		response := server_mock.NewHttpResponseWriterMock()
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(response, request)
		err = AssertThat(response.String(), Is(content).Message("check content"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/api/test?paramName=paramValue")
		if err != nil {
			t.Error(err)
		}
		response := server_mock.NewHttpResponseWriterMock()
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(response, request)
		err = AssertThat(response.String(), Is(content).Message("check content"))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestWithQuestionMarkWithoutSlash(t *testing.T) {
	content := "TestHandlerContent"
	hf := New("/hello")
	hf.RegisterHandler("world", static.New(content))
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/helloworld")
		if err != nil {
			t.Error(err)
		}
		response := server_mock.NewHttpResponseWriterMock()
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(response, request)
		err = AssertThat(response.String(), Is(content).Message("check content"))
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		request, err := server_mock.NewHttpRequestMock("http://www.example.com/helloworld?paramName=paramValue")
		if err != nil {
			t.Error(err)
		}
		response := server_mock.NewHttpResponseWriterMock()
		handler := hf.FindHandler(request)
		err = AssertThat(handler, NotNilValue().Message("handler registered"))
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(response, request)
		err = AssertThat(response.String(), Is(content).Message("check content"))
		if err != nil {
			t.Fatal(err)
		}
	}
}
