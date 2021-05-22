package services

import (
	"fmt"
	dto "go-auth/dto/login"
	"go-auth/repositories/user"
	"go-auth/services"
)

type LoginService struct {
	repo user.UserRepository
}

func (s *LoginService) GetToken(dto dto.LoginRequest) (string, error) {
	user, err := s.repo.GetByMail(dto.Mail)
	if err != nil {
		return "", err
	}

	pass, err := services.SHA256Encoder(dto.Pass)
	if err != nil {
		return "", err
	}

	if user.Password == pass {
		token := services.NewJWTService().GenerateToken(user.Name, dto.Mail)

		return token, nil
	}

	return "", fmt.Errorf("credenciais inválidas")
}

func NewLoginService(r user.UserRepository) *LoginService {
	return &LoginService{
		repo: r,
	}
}
