package main

import "net/http"

type AuthMiddleware struct {
	Handler http.Handler
	User string
	Pass string
}

func (auth *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	user, pass, ok := request.BasicAuth()

	if !ok {
		writer.Header().Add("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(writer, "No basic auth present", http.StatusUnauthorized)
		return
	}

	if !(user == auth.User && pass == auth.Pass) {
		writer.Header().Add("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(writer, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	auth.Handler.ServeHTTP(writer, request)
}
