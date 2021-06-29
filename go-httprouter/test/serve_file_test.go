package test

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeFile(t *testing.T)  {
	router := httprouter.New()
	router.ServeFiles("/files/*filepath", http.Dir("./resources"))

	request := httptest.NewRequest("GET", "http://localhost:9191/files/data.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Data1", string(body))
}

//go:embed resources
var resources embed.FS

func TestServeFileEmbed(t *testing.T)  {
	router := httprouter.New()
	dir, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest("GET", "http://localhost:9191/files/data.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Data1", string(body))
}
