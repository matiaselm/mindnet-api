package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
