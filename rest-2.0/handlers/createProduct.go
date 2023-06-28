package handlers

import (
	"encoding/json"
	"home/geetha/Desktop/practice/rest-2.0/database"
	"home/geetha/Desktop/practice/rest-2.0/models"
	"net/http"
)

func CreateProduct(rw http.ResponseWriter, r *http.Request) {

	// add product
	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)

	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
		return
	}

	//recieve return types and send response code
	_, err1 := database.CreateProduct(*product)
	if err1 != nil {
		http.Error(rw, "Unable to create", http.StatusBadRequest)
		return
	}

}
