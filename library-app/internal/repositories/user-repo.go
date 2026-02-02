package repositories

import (
	"database/sql"
	"errors"                      // para crear errores personalizados
	"library-app/internal/models" // para  usar objeto User
)

// estructura se encarga únicamente de acceder a la tabla users

type UserRepository struct {
	db *sql.DB // conexión a postgresql
}

// constructor del repositorio
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

//	busca un usuario por su email en la base de datos
//
// Ejecuta un SELECT filtrando por email
//
//	Si existe, mapea la fila a un struct user
//
// retorna objeto user, y mensaje de error
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {

	// estructura donde se guarda el usuario
	var user models.User

	err := r.db.QueryRow(
		"SELECT id, email, role FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Email, &user.Role)

	switch err {

	case sql.ErrNoRows:
		return nil, errors.New("usuario no encontrado")
	case nil:
		return &user, nil
	default:
		return nil, err
	}

}
