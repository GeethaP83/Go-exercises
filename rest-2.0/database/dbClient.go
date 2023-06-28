package database

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"

	"home/geetha/Desktop/practice/rest-2.0/models"
)

type Database interface {
	GetAllProducts() ([]models.Product, error)
	GetProductById(id int) (models.Product, error)
	CreateProduct(product models.Product) (bool, error)
	// GetReviewsByProductId(id int) ([]Review, error)
	// CreateReviewForProductId(review Review, id int) (bool, error)
}

var DB *gorm.DB

func SetupDB() {
	host := flag.String("host", "localhost", "PostgreSQL host")
	port := flag.String("port", "5432", "PostgreSQL port")
	user := flag.String("user", "postgres", "PostgreSQL user")
	password := flag.String("password", "password", "PostgreSQL password")
	database := flag.String("database", "postgres", "PostgreSQL database name")
	flag.Parse()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", *host, *port, *user, *password, *database)
	db, err := gorm.Open("postgres", connectionString)
	checkErr(err)
	DB = db
	//return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
