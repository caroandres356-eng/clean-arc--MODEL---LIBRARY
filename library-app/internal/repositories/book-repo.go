package repositories

import (
	"database/sql"
	"library-app/internal/models" // para trabajar con objwtos libro
)

//	estructura ue se comunica con la db
//
// solo tiene consultas sql
// NO contiene reglas de negocio
type BookRepository struct {
	db *sql.DB // conexión activa a postgresql
}

// inicialzia estructura que recibe una conexión ya creada
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

//  obtiene todos los libros desde la tabla "books"

//	Ejecuta la consulta SQL con db.Query
//	PostgreSQL devuelve filas
//
// Recorremos cada fila con rows.Next
// Convertimos cada fila en un struct Book
// Guardamos los libros en un slice
// Retornamos el slice completo
//
// retorna slice y mensaje de error
func (r *BookRepository) GetAll() ([]models.Book, error) {

	// ejecuta la consulta en la base de datos
	rows, err := r.db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// slice donde guardaremos los libros
	var books []models.Book

	// recorremos cada fila
	for rows.Next() {

		// estructura temporal para guardar un libro
		var b models.Book

		// copiamos los valores de la fila en la estructura Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}

		// agregamos el libro al slice
		books = append(books, b)
	}

	// devolvemos todos los libros encontrados
	return books, nil
}

//  inserta un nuevo libro en la base de datos

//  Ejecuta un INSERT usando db.Exec
// la db guarda el registro

// recibe el libro que va a insertar
// Retorna mensaje de error
func (r *BookRepository) Add(book models.Book) error {

	// ejecutamos el INSERT con parámetros para evitar SQL injection
	_, err := r.db.Exec(
		"INSERT INTO books (title, author) VALUES ($1, $2)",
		book.Title,
		book.Author,
	)

	return err
}
