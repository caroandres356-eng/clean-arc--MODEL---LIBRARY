package middlewares

import (
	"context"
	"net/http"
	"strings"

	"library-app/internal/auth"
)

type contextKey string

const (
	roleKey   contextKey = "role"
	userIDKey contextKey = "userID"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization requerido", http.StatusUnauthorized)
			return
		}

		// Esperamos: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "formato de token inválido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "token inválido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		ctx = context.WithValue(ctx, roleKey, claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
