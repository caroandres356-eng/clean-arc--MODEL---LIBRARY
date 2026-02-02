package services

// aqui va la logica de negocio
import "library-app/internal/repositories" //se importan los repositorios para operar con el repositorio de usuarios

type AuthService struct {
	UserRepo *repositories.UserRepository
} //  estructura de datos que contiene la estructura que contiene los metodos de acceso a db

func (s *AuthService) Login(email string) (string, error) {

	// buscamos el usuario en el repositorio
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	// generamos token que es el rol de quien hace la peticion
	token := "token-" + user.Role

	return token, nil
}
