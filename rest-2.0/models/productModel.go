package models

type Product struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Category    string
	Quantity    uint
	Price       float32
	Image       string
	Reviews     []Review
}
