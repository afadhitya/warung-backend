package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Good    Good
	GoodID  uint
	Amount  int
	OrderID uint
	Order   Order
}
