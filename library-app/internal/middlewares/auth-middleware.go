package middlewares

import (
	"context" //para asignar role despues de que pase por el middleware
	"net/http"
	"strings"
)

// identifica el rol de cada peticion
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")             //almacena el token de la peticion
		role := strings.TrimPrefix(token, "Bearer token-") //almacena elrol de cada peticion

		ctx := context.WithValue(r.Context(), "role", role) //asigna el rol a el request
		next.ServeHTTP(w, r.WithContext(ctx))               //sigue al siguiente handler le envia la respuesta del handler y el contexto con el  rol
	})
}
