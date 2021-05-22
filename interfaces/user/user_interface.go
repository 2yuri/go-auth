package user

import (
	"go-auth/models"
)

type UserRepository interface {
	Get(id uint) (models.User, error)
	Create(dto models.User) (models.User, error)
	Edit(dto models.User) error
	Delete(id uint) error
	GetByMail(mail string) (models.User, error)
}
