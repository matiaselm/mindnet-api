package models
/*
import (
   "gorm.io/gorm"
)

type Technology struct {
    gorm.Model
    RaceId       int
    UnitId       int
    Name         string
    Type         string
    Effect       string
    Requirements string
}

//create a technology
func CreateTechnology(db *gorm.DB, Technology *Technology) (err error) {
   err = db.Create(Technology).Error
   if err != nil {
      return err
   }
   return nil
}

//get technologies
func GetTechnologies(db *gorm.DB, Technology *[]Technology) (err error) {
   err = db.Find(Technology).Error
   if err != nil {
      return err
   }
   return nil
}

//get technology by id
func GetTechnology(db *gorm.DB, Technology *Technology, id string) (err error) {
   err = db.Where("id = ?", id).First(Technology).Error
   if err != nil {
      return err
   }
   return nil
}

//update technology
func UpdateTechnology(db *gorm.DB, Technology *Technology) (err error) {
   db.Save(User)
   return nil
}

//delete technology
func DeleteTechnology(db *gorm.DB, Technology *Technology, id string) (err error) {
   db.Where("id = ?", id).Delete(User)
   return nil
}*/