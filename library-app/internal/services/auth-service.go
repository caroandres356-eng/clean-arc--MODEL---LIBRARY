package services

import (
	"errors"

	"library-app/internal/auth"
	"library-app/internal/repositories"
	"library-app/internal/utils"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func (s *AuthService) Login(email, password string) (string, error) {

	// 1️⃣ buscar usuario
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("credenciales inválidas")
	}

	// 2️⃣ comparar password (bcrypt)
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	// 3️⃣ generar JWT
	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(email, password string) error {

	// hash password
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// guardar usuario
	return s.UserRepo.Create(email, hash)
}
