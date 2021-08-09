package models

import (
	"gorm.io/gorm"
)

type Race struct {
	gorm.Model
	ID     int
	Name   string `json:"name" binding:"required"`
	Banner string `json:"banner" binding:"required"`
}
