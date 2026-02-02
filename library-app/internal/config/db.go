package config

import (
	"database/sql" //paquete estandar para trabajar con sql db

	_ "github.com/lib/pq" //driver entre go y postgre
)

// retorna estrutura manejadora de conexiones(pool de conexiones) a la db y mensaje de error
func ConnectPostgres() (*sql.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=1014736507 dbname=library_db sslmode=disable" //data source name
	return sql.Open("postgres", dsn)
}
