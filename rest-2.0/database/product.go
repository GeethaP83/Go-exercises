package database

import (
	"fmt"

	"home/geetha/Desktop/practice/rest-2.0/models"
)

// GetProducts returns the list of Products
func GetAllProducts() ([]models.Product, error) {

	fmt.Println("Getting products...")

	results := []models.Product{}
	err := DB.Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return results, err
}

// GetProductById returns a unique Product
func GetProductById(id int) (models.Product, error) {

	fmt.Println("Getting product by ID ")
	result := models.Product{
		ID: id,
	}
	err := DB.Where(&result).Find(&result).Error
	if err != nil {
		return models.Product{}, fmt.Errorf(err.Error())
	}
	return result, err
}

// AddProduct adds a Product in the DB
func CreateProduct(product models.Product) (bool, error) {

	fmt.Println("Adding new product")

	err := DB.Create(&product).Error
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	return true, err

}
