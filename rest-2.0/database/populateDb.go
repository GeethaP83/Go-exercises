package database

import (
	"github.com/jinzhu/gorm"
	"home/geetha/Desktop/practice/rest-2.0/models"
)

func InsertData(db *gorm.DB) error {

	products := []models.Product{
		{Name: "NameA", Description: "DescriptionA", Category: "CategoryA", Quantity: 2, Price: 4, Image: "ImageA"},
		{Name: "NameB", Description: "DescriptionB", Category: "CategoryB", Quantity: 3, Price: 6, Image: "ImageB"},
	}

	for _, product := range products {
		result := db.Create(&product)
		if result.Error != nil {
			return result.Error
		}
	}

	reviews := []models.Review{
		{Name: "Reviewer1", Comment: "Comment1A", Rating: 3, ProductId: 1},
		{Name: "Reviewer2", Comment: "Comment2A", Rating: 4, ProductId: 2},
		{Name: "Reviewer1", Comment: "Comment1B", Rating: 1, ProductId: 1},
		{Name: "Reviewer2", Comment: "Comment2B", Rating: 2, ProductId: 2},
	}

	for _, review := range reviews {
		result := db.Create(&review)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
