package config

import (
	"fmt"
	"log"
	"os"

	"cihanozhan.com/dbgo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	DB *gorm.DB
)

func ConnDB() *gorm.DB {
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	if dbHost == "" {
		dbHost = "localhost"
	}

	if dbUser == "" {
		dbUser = "postgres"
	}

	if dbPassword == "" {
		dbPassword = "Nurcan1234"
	}

	if dbName == "" {
		dbName = "usersdb"
	}

	log.Println("DB --> " + "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost)
	// db, err := gorm.Open("postgres", "user="+dbUser+" password="+dbPassword+" dbname="+dbName+" sslmode=disabled")
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

	return db
}
