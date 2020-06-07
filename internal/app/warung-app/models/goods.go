package models

import "github.com/jinzhu/gorm"

type Good struct {
	gorm.Model
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	Stock      int64    `json:"stock"`
	Discount   float64  `json:"discount"`
	Category   Category `json:"category"`
	CategoryID uint     `json:"category-id"`
}
