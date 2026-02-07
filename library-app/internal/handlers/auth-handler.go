package handlers

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

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	if err := h.Service.Register(body.Email, body.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login recibe una petición de login,
// extrae los datos necesarios,
// llama al servicio de autenticación
// y devuelve un token si la autenticación es correcta
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	// estructura local para leer los datos del body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// decodificamos el json del cuerpo que viene desde la peticion
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	// delegamos la lógica de login al servicio
	token, err := h.Service.Login(body.Email, body.Password)

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
