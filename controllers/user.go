package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"gorm-mysql/database"
	"gorm-mysql/models"
)

type UserRepo struct {
	Db *gorm.DB
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	
	c.JSON(http.StatusOK, gin.H{"data": users})
}

//create user
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
	}

	// Create book
	user := models.User{Name: input.Name, Email: input.Email}
	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}