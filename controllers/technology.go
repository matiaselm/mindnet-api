package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"mindnet-api/database"
	"mindnet-api/models"
)

type TechnologyRepo struct {
	Db *gorm.DB
}

type CreateTechnologyInput struct {
	RaceId         int
	UnitId         int
	Name           string `json:"name" binding:"required"`
	Class          string `json:"class" binding:"required"`
	Effect         string `json:"effect" binding:"required"`
	Requirements   string `json:"requirements" binding:"required"`
}

func GetTechnologies(c *gin.Context) {
	var technologies []models.Technology
	database.DB.Find(&technologies)
	
	c.JSON(http.StatusOK, gin.H{"data": technologies})
}

//create technology
func CreateTechnology(c *gin.Context) {
	// Validate input
	var input CreateTechnologyInput
	if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	technology := models.Technology{RaceId: input.RaceId, UnitId: input.UnitId, UnitName: input.Name, Class: input.Class, Effect: input.Effect, Requirements: input.Requirements}
	database.DB.Create(&technology)

	c.JSON(http.StatusOK, gin.H{"data": technology})
}