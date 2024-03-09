package middlewares

import (
	"log"
	"net/http"
)

// Middleware that prints "middleware"
func PrintMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware")
		next.ServeHTTP(w, r)
	})
}
