package models

type Role struct {
	ID    uint   `gorm:"primaryKey"`
	Label string `json:"label"`
}
