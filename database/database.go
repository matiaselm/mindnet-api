package database

import (
	"errors"
	"mindnet-api/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "admin"
const DB_PASSWORD = "admin"
const DB_NAME = "ti_assist"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var DB *gorm.DB


// AutoMigrateDB will keep tables reflecting structs
func AutoMigrateDB(DB *gorm.DB) error {
	// if tables exists check if they reflects struts
	if err := DB.AutoMigrate(models.&User, models.&Technology, models.&Unit, models.&Race).Error; err != nil {
		return errors.New("Unable autoMigrateDB - " + err.Error())
	}
	return nil
}

func ConnectDB() (*gorm.DB) {
	var err error
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	DB = db
	AutoMigrateDB(DB)
	return db
}
