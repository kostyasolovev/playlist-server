package app

import (
	"log"
	"net/http"
)

// отлавливаем панику.
func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic recovered:", err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
