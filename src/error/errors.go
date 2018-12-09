package error

import (
	"net/http"

	u "../utils"
)

//NotFoundHandler error needs to be thrown in case any error occurs in the server
var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Resource is not found in the server"))
		next.ServeHTTP(w, r)
	})
}
