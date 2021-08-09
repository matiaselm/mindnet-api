package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"mindnet-api/database"
	"mindnet-api/controllers"
)

func main() {
	r := gin.Default()
	
	database.ConnectDB()
	// database.AutoMigrateDB()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
	})
	
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	
	r.Run()
}