package config

import (
	_ "database/sql"

	"github.com/afadhitya/warung-backend/internal/app/warung-app/models"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {

	dbConnectionStr := "golang-user@/warung-db?charset=utf8&parseTime=True&loc=Local"
	// "afadhitya:password@cloudsql(fluted-insight-279501:asia-southeast2:warung-adita-db)/warung_db?charset=utf8&parseTime=True&loc=UTC"
	var err error
	DB, err = gorm.Open("mysql", dbConnectionStr)
	if err != nil {
		panic("failed to connect database")
	}

	DB.DropTableIfExists(
		&models.Category{},
		&models.Good{},
		&models.Item{},
		&models.User{},
		&models.Order{},
	)
	DB.AutoMigrate(
		&models.Category{},
		&models.Good{},
		&models.Item{},
		&models.User{},
		&models.Order{},
	)

	// createDummyData()
}

func createDummyData() {
	Cat1 := models.Category{
		Name:       "drink",
		CodeNumber: "1234",
		Goods: []models.Good{
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
