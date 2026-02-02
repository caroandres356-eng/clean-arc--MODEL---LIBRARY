package services

import (
	"errors"
	"library-app/internal/models"       // traemos los modelos para trabajar con libros
	"library-app/internal/repositories" // traemos el repositorio que habla con la DB
)

// estructura que almacena la estructura que se comunica con la db
type BookService struct {
	Repo *repositories.BookRepository
}

//	retorna todos los libros desde la base de datos
//
// llama al repositorio y maneja posibles errores
func (s *BookService) GetBooks() ([]models.Book, error) {

	// pedimos los libros al repositorio
	books, err := s.Repo.GetAll()

	// si algo falla propagamos el error
	if err != nil {
		return nil, err
	}
	return books, nil
}

// agrega un libro a la base de datos
// recibe el libro y el rol del usuario para validar permisos
func (s *BookService) AddBook(book models.Book, role string) error {

	if role != "admin" {
		return errors.New("solo admin puede agregar libros")
	}

	// guardamos el libro usando el repositorio
	return s.Repo.Add(book)
}
