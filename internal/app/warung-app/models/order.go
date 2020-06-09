package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	User  User
	Items []Item
	Total float64
}
