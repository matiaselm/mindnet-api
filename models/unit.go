package models

import (
	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	ID     int
	Name   string `json:"name" binding:"required"`
	Level  string `json:"level"`
}
