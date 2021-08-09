package models

import (
	"gorm.io/gorm"
)

type Technology struct {
	gorm.Model
	ID             int
	Race_id        int
   Unit_id        int
   Name           string `json:"name" binding:"required"`
   Class          string `json:"class" binding:"required"`
   Effect         string `json:"effect" binding:"required"`
   Requirements   string `json:"requirements" binding:"required"`
}