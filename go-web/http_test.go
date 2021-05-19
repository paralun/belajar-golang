package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(writer, "Hallo Test")
}

func HelloQueryParam(writer http.ResponseWriter, request *http.Request)  {
	firstname := request.URL.Query().Get("firstname")
	lastname := request.URL.Query().Get("lastname")
	if firstname == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
	}
}

func TestHttp(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestHttpQueryParam(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000/hello?firstname=paralun&lastname=keren", nil)
	recorder := httptest.NewRecorder()

	HelloQueryParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
