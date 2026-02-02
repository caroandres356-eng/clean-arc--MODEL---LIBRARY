package config

// contiene la informacion global de donde correra el servidor , pagina web etc..
type Config struct {
	Port string
}

// carga la infraestructura , es cmo un metodo estatico
func Load() Config {
	return Config{
		Port: ":8080",
	}
}
