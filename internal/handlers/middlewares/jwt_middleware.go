package middlewares

import (
	"net/http"
	"project-control-system/pkg/jwt_utils"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "[JWTMiddleware]: Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "[JWTMiddleware]: Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		isTokenValid := jwt_utils.ValidateToken(tokenString)
		if !isTokenValid {
			http.Error(w, "[JWTMiddleware]: Invalid token", http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r)
	})
}
