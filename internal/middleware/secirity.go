package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}).Handler(next)
}

func LogMiddlewereAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\033[1;33m%s\033[0m \033[1;34m%s\033[0m Access log: %s \033[1;32m#\033[0m %s \033[1;31m@\033[0m %s\n", time.Now().Format(time.RFC3339), time.Now().Format(time.Kitchen), r.Method, r.URL.Path, r.RemoteAddr)
		w.Header().Set("X-Hello", "World")
		// Do something before the request
		next.ServeHTTP(w, r)
		// Do something after the request
	})
}
