package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"home/geetha/Desktop/practice/rest-2.0/database"
	"home/geetha/Desktop/practice/rest-2.0/handlers"
	"home/geetha/Desktop/practice/rest-2.0/models"
)

func main() {
	database.SetupDB()
	defer database.DB.Close()
	fmt.Println("Connection Made")

	database.DB.DropTable(&models.Product{}, &models.Review{})
	database.DB.AutoMigrate(&models.Product{}, &models.Review{})
	if err := database.InsertData(database.DB); err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	//handle routes
	router.HandleFunc("/products/", handlers.GetAllProducts).Methods("GET")

	router.HandleFunc("/products/{id}", handlers.GetProductById).Methods("GET")

	router.HandleFunc("/products/create", handlers.CreateProduct).Methods("POST")

	// router.HandleFunc("/products/{id}/reviews", handlers.GetReviewsForProduct).Methods("GET")

	// router.HandleFunc("/products/{id}/reviews/create", handlers.CreateReviewForProductId).Methods("POST")

	log.Fatal(http.ListenAndServe(":8020", router))
}
