package middleware

import (
	"log"
	"net/http"
)

// BasicLoggingHandler
func BasicLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(*r.URL)
		next.ServeHTTP(w, r)
	})
}
