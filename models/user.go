package models

import (
	"go-auth/services"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	JWTToken  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewUser(name string, email string, password string) *User {
	if name == "" || email == "" || password == "" {
		return nil
	}

	pass, err := services.SHA256Encoder(password)
	if err != nil {
		return nil
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: pass,
	}
}

func NewUserWithId(id uint, name string, email string, encodedPass string) *User {
	if id == 0 || name == "" || email == "" || encodedPass == "" {
		return nil
	}
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: encodedPass,
	}
}
