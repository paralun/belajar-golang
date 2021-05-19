package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPostParseManual(writer http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	first := request.PostForm.Get("firstname")
	last := request.PostForm.Get("lastname")

	fmt.Fprintf(writer, "Hallo %s %s", first, last)
}

func FormPostParseAuto(writer http.ResponseWriter, request *http.Request)  {

	first := request.PostFormValue("firstname")
	last := request.PostFormValue("lastname")

	fmt.Fprintf(writer, "Hallo %s %s", first, last)
}

func TestRequestPostForm(t *testing.T)  {
	reqBody := strings.NewReader("firstname=paralun&lastname=ganteng")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:7000/", reqBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	//FormPostParseManual(recorder, request)
	FormPostParseAuto(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}