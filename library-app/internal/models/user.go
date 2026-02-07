package models

type User struct {
	ID       int
	Email    string
	Password string // HASH, no texto plano
	Role     string
}
