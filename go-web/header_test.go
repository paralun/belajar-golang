package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request)  {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Add("X-Paralun-Key", "pas123456")
	fmt.Fprint(writer, "OK")
}

func TestRequestHeader(t *testing.T)  {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:7000/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:7000/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	resHeader := response.Header.Get("x-paralun-key")
	fmt.Println(resHeader)
}