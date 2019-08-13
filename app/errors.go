package app

import (
	"net/http"

	"github.com/jhindulak/go-rest-api-example/utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		utils.Respond(w, utils.Message(false, "This resource was not found on the server"))
		next.ServeHTTP(w, r)
	})
}
