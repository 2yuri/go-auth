package migrations

import (
	"go-auth/models"

	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
