package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`

	Todos []Todo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"todos"`
}
