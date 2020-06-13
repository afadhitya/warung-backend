package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	User   User
	UserID int
	Items   []Item
	Total  float64
	Status int
}

const STATUS_ON_CART = 1
const STATUS_WAITING_FOR_PAYMENT = 2
const STATUS_PAID = 3
const STATUS_DONE = 4
