package config

import (
	"fmt"
	"log"

	"cihanozhan.com/dbgo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	DB *gorm.DB
)

func ConnDB() {
	var err error
	db, err := gorm.Open("postgres", "user=postgres password=Nurcan1234 dbname=usersdb sslmode=disable")

	fmt.Println(err)
	if err != nil {
		log.Println("Database connection errror", err)
		panic(err)
	}
	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Database connection established")
	fmt.Println("Successfuly connected to the database.")

}
