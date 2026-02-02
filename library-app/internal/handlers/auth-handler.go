package handlers

// NO contiene lógica de autenticación ni autorización.
// - recibir la peticion http
// - decodificar el cuerpo del request
// - delegar la lógica al servicio
// - devolver la respuesta http

import (
	"encoding/json"
	"library-app/internal/services" // servicio que contiene la lógica de autenticación
	"net/http"
)

// AuthHandler, estrucutra de datos que almacena el servicio de autenticación
// el handler solo coordina la entrada y salida http
type AuthHandler struct {
	Service *services.AuthService
}

// Login recibe una petición de login,
// extrae los datos necesarios,
// llama al servicio de autenticación
// y devuelve un token si la autenticación es correcta
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	// estructura local para leer los datos del body
	var body struct {
		Email string `json:"email"`
	}

	// decodificamos el json del cuerpo que viene desde la peticion
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	// delegamos la lógica de login al servicio
	token, err := h.Service.Login(body.Email) //el token es el rol
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// devolemos un objeto json como respuesta http que tiene el vslor de token(que es el rol)
	w.Header().Set("Content-Type", "application/json")

	// devolvemos el token obtenido
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
