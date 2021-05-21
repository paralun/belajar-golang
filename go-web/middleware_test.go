package go_web

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type LogHandler struct {
	Handler http.Handler
}

func (logHandler *LogHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	log.Println("Before Execute Handler")
	logHandler.Handler.ServeHTTP(writer, request)
	log.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Middleware")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("Error euy")
	})

	logHandler := &LogHandler{Handler:mux}
	errorHandler :=&ErrorHandler{Handler:logHandler}

	server := http.Server{
		Addr: "localhost:7000",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}