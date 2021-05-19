package models

import (
	"fmt"
	"go-auth/services"
	"time"

	"gorm.io/gorm"
)

type User struct {
	id        uint `gorm:"primaryKey"`
	name      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
	deletedAt gorm.DeletedAt `gorm:"index"`
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
		name:     name,
		email:    email,
		password: pass,
	}
}

func NewUserWithId(id uint, name string, email string, encodedPass string) *User {
	if id == 0 || name == "" || email == "" || encodedPass == "" {
		return nil
	}
	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: encodedPass,
	}
}

func (s *User) GetId() uint {
	return s.id
}

func (s *User) GetName() string {
	return s.name
}

func (s *User) GetEmail() string {
	return s.email
}

func (s *User) GetPassword() string {
	return s.password
}

func (s *User) String() string {
	return fmt.Sprintf("{%v %v %v}", s.name, s.email, s.password)
}
