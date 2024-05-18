package models

import (
	"go-crud/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Movies struct {
	gorm.Model  //this adds default ID, createdAt fields
	Name        string
	HeroName    string
	HeroineName string
}

func init() {
	//connect to database

	config.ConnectToDB()

	db = config.GetDb()

	//auto increment add the code later
	db.AutoMigrate(&Movies{})

}

func (m *Movies) CreateMovie() *Movies {

	//ceate new Record
	db.NewRecord(m)

	//insert into database
	db.Create(&m)

	//return the movie just created
	return m

}
