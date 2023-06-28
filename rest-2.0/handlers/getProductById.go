package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"home/geetha/Desktop/practice/rest-2.0/database"
	"home/geetha/Desktop/practice/rest-2.0/models"
	"net/http"
	"strconv"
)

func GetProductById(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	// fetch the product
	product := &models.Product{}
	*product, err = database.GetProductById(id)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}
	err1 := json.NewEncoder(rw).Encode(*product)
	if err1 != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}
}
