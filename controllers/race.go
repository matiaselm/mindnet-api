package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"mindnet-api/database"
	"mindnet-api/models"
)

type RaceRepo struct {
	Db *gorm.DB
}

type CreateRaceInput struct {
	Name  string `json:"name" binding:"required"`
	Banner string `json:"banner" binding:"required"`
}

func GetRaces(c *gin.Context) {
	var races []models.Race
	database.DB.Find(&races)
	
	c.JSON(http.StatusOK, gin.H{"data": races})
}

// Create race
func CreateRace(c *gin.Context) {
	// Validate input
	var input CreateRaceInput
	if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	race := models.Race{Name: input.Name, Banner: input.Banner}
	database.DB.Create(&race)

	c.JSON(http.StatusOK, gin.H{"data": race})
}