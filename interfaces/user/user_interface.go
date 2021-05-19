package user

import (
	"go-auth/dto/user"
	"go-auth/models"
)

type UserRepository interface {
	Get(id uint) (*models.User, error)
	Create(dto user.UserNewDTO) (*models.User, error)
	Edit(dto user.UserEditDTO) error
	Delete(id uint) error
}
