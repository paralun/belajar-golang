package test

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterParam(t *testing.T)  {
	router := httprouter.New()
	router.GET("/hello/:name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		name := params.ByName("name")
		text := "Hello " + name
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:9191/hello/dunia", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello dunia", string(body))
}
