package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name       string `json:"name"`
	CodeNumber string `json:"code_number"`
	Goods      []Good `json:"goods"`
}
