package middleware

import (
	"log"
	"net/http"
)

func (m *Middleware) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle token here
		log.Println("JWT is called!")
		next.ServeHTTP(w, r)
	})
}
