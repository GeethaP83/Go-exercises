package database

import (
	"fmt"
	"home/geetha/Desktop/practice/rest-2.0/models"
)

// GetReviewsById returns review of a particular product
func GetReviewsByProductId(id int) ([]models.Review, error) {

	fmt.Println("Getting reviews of Product with ID ")

	reviews := []models.Review{}

	err := DB.Where(models.Review{ProductId: id}).Find(&reviews).Error

	return reviews, err
}

// AddReview adds a review for a particular product
func CreateReviewForProductId(review models.Review, id int) (bool, error) {

	fmt.Println("Adding new review")
	review.ProductId = id
	err := DB.Save(&review).Error
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	return true, err
}
