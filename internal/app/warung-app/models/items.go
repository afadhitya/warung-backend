package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Good    Good
	Amount  int
	OrderID uint
	Order   Order
}
