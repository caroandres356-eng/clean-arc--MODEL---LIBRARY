package handlers

import (
	"encoding/json"
	"net/http"

	"library-app/internal/models"   // importa los modelos (estructuras de datos)
	"library-app/internal/services" // importa los servicios para usar sus metodos
)

// BookHandler, estructura de datps que almacena el servicio de libros
type BookHandler struct {
	Service *services.BookService
}

// Lista los libros y los envía como archivo json
func (h *BookHandler) List(w http.ResponseWriter, r *http.Request) {

	// llamamos al metodo del servico para obtener los libros
	books, err := h.Service.GetBooks()

	if err != nil {
		http.Error(w, "error obteniendo libros", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// enviamos los libros en  json
	json.NewEncoder(w).Encode(books)
}

// agrega un libro nuevo
func (h *BookHandler) Add(w http.ResponseWriter, r *http.Request) {

	// obtenemos el rol desde el contexto
	role := r.Context().Value("role").(string)

	// creamos una estructura libro vacía
	var book models.Book

	// decodificamos el json del body de la petición en la estructura
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "json invalido", http.StatusBadRequest)
		return
	}

	// llamamos al servicio para agregar el libro
	if err := h.Service.AddBook(book, role); err != nil {

		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
