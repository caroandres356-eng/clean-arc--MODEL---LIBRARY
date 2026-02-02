package main

import (
	"fmt"
	"net/http"

	"library-app/internal/config"
	"library-app/internal/handlers"
	"library-app/internal/repositories"
	routes "library-app/internal/routing"
	"library-app/internal/services"
)

func main() {

	cfg := config.Load()
	mux := http.NewServeMux()

	// crear la conexion
	db, err := config.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	//repositorios
	userRepo := repositories.NewUserRepository(db)
	bookRepo := repositories.NewBookRepository(db)

	// servicios
	authService := &services.AuthService{UserRepo: userRepo}
	bookService := &services.BookService{Repo: bookRepo}

	// handlers
	authHandler := &handlers.AuthHandler{Service: authService}
	bookHandler := &handlers.BookHandler{Service: bookService}

	// rutas
	routes.Register(mux, authHandler, bookHandler)

	fmt.Println("servidor corriendo en el uerto:", cfg.Port)
	http.ListenAndServe(cfg.Port, mux)
}
