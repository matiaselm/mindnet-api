package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"mindnet-api/database"
	"mindnet-api/models"
)

type UnitRepo struct {
	Db *gorm.DB
}

type CreateUnitInput struct {
	Name  string `json:"name" binding:"required"`
	Level string `json:"level"`
}

func GetRaces(c *gin.Context) {
	var units []models.Unit
	database.DB.Find(&units)
	
	c.JSON(http.StatusOK, gin.H{"data": units})
}

// Create unit
func CreateUnit(c *gin.Context) {
	// Validate input
	var input CreateUnitInput
	if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	unit := models.Unit{Name: input.Name, Level: input.Level}
	database.DB.Create(&unit)

	c.JSON(http.StatusOK, gin.H{"data": unit})
}