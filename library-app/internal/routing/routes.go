package routes

// enrutamiento de handlers
import (
	"net/http"

	"library-app/internal/handlers"    // importamos paquetes de handlers
	"library-app/internal/middlewares" // importamos paquete de middlewares
)

func Register(mux *http.ServeMux, auth *handlers.AuthHandler, book *handlers.BookHandler) { // multiplexer para asignar rutas  ,  handler de logeo principal,  handlers de libro(para agregr y listar libros)

	mux.HandleFunc("/login", auth.Login) //ruta inicial de logeo

	mux.Handle("/books",
		middlewares.Auth(http.HandlerFunc(book.List)),
	) //para acceder a la lista de libros , primero debe tene el cotnexto asignado por el middleware con el rol

	mux.Handle("/books/add",
		middlewares.Auth(http.HandlerFunc(book.Add)),
	) //para añadir un libro primero necesito el contexto con el rol nuevamente
}
