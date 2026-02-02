package models

type User struct {
	ID    int
	Email string
	Role  string // "admin" o "user"
}
