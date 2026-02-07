package routes

import (
	"net/http"

	"library-app/internal/handlers"
	"library-app/internal/middlewares"
)

func Register(
	mux *http.ServeMux,
	auth *handlers.AuthHandler,
	book *handlers.BookHandler,
) {

	// 🔓 RUTAS PÚBLICAS
	mux.HandleFunc("/login", auth.Login)
	mux.HandleFunc("/register", auth.Register)

	// 🔒 RUTAS PROTEGIDAS
	mux.Handle(
		"/books",
		middlewares.Auth(http.HandlerFunc(book.List)),
	)

	mux.Handle(
		"/books/add",
		middlewares.Auth(http.HandlerFunc(book.Add)),
	)
}
