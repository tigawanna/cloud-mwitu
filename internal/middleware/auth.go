package middleware

import (

	"net/http"

	"github.com/tigawanna/cloud-mwitu/internal/controllers"
	"github.com/tigawanna/cloud-mwitu/internal/services"
)



func AuthMiddleware(authService services.AuthService) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Get session cookie
            cookie, err := r.Cookie(controller.SessionCookieName)
            if err != nil {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            sessionID := cookie.Value
            
            // Validate the session
            _, err = authService.ValidateSession(sessionID)
            if err != nil {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            // Call the next handler
            next.ServeHTTP(w, r)
        })
    }
}

