package handlers

import (
	"encoding/json"
	"home/geetha/Desktop/practice/rest-2.0/database"
	"net/http"
)

func GetAllProducts(rw http.ResponseWriter, r *http.Request) {

	// fetch products

	products, err := database.GetAllProducts()
	if err != nil {
		http.Error(rw, "Products not found", http.StatusBadRequest)
		return
	}

	err1 := json.NewEncoder(rw).Encode(products)

	if err1 != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}
}
