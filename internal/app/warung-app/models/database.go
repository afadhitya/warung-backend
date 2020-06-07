package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root@/warung_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	DB.DropTableIfExists(&Category{}, &Good{})
	DB.AutoMigrate(&Category{}, &Good{})

	createDummyData()
}

func createDummyData() {
	Cat1 := Category{
		Name:       "drink",
		CodeNumber: "1234",
		Goods: []Good{
			{
				Name:     "prutang",
				Price:    7000,
				Stock:    23,
				Discount: 0.5,
			},
			{
				Name:     "panta",
				Price:    7000,
				Stock:    23,
				Discount: 0.4,
			},
		},
	}

	DB.Create(&Cat1)
}