package models

import (
	"github.com/jinzhu/gorm"
)

type Review struct {
	gorm.Model
	Name      string
	Comment   string
	Rating    uint
	ProductId int
}
