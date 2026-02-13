package auth

import (
	"errors" // Permite crear y manejar errores personalizados
	"os"     // Permite acceder a variables de entorno
	"time"   // Permite manejar fechas y tiempos (expiración del token)

	// Librería para crear y validar JSON Web Tokens (JWT)
	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret almacena la clave secreta usada para firmar y validar los tokens JWT.
// Se obtiene desde la variable de entorno JWT_SECRET por seguridad.
// Nunca debe estar hardcodeada en el código.
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Claims define la información (payload) que se almacenará dentro del token JWT.
type Claims struct {
	// ID del usuario autenticado
	UserID int `json:"user_id"`

	// Rol del usuario (ej: "admin", "user")
	Role string `json:"role"`

	// Claims estándar del JWT (expiración, fecha de emisión, etc.)
	jwt.RegisteredClaims
}

// GenerateToken genera un token JWT firmado para un usuario.
// Recibe:
// - userID: ID del usuario
// - role: rol del usuario
// Retorna:
// - string: el token JWT
// - error: error si ocurre algún problema al firmar el token
func GenerateToken(userID int, role string) (string, error) {

	// Se crean los claims (datos) que irán dentro del token
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			// El token expirará 1 hora después de su creación
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	// Se crea un nuevo token usando el algoritmo HS256 y los claims definidos
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Se firma el token usando la clave secreta y se devuelve como string
	return token.SignedString(jwtSecret)
}

// ValidateToken valida un token JWT recibido como string.
// Recibe:
// - tokenString: el token JWT enviado por el cliente
// Retorna:
// - *Claims: los datos contenidos en el token si es válido
// - error: error si el token es inválido o está expirado
func ValidateToken(tokenString string) (*Claims, error) {

	// Se parsea y valida el token, indicando qué estructura usar para los claims
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		// Función que devuelve la clave secreta para verificar la firma del token
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)

	// Si ocurre un error o el token no es válido, se retorna error
	if err != nil || !token.Valid {
		return nil, errors.New("token inválido")
	}

	// Se convierten los claims del token al tipo Claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("claims inválidos")
	}

	// Se devuelven los claims del token
	return claims, nil
}
