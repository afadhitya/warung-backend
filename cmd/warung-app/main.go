package main

import (
	"fmt"

	"github.com/afadhitya/warung-app/internal/app/warung-app/models"
	"github.com/afadhitya/warung-app/internal/app/warung-app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("application start")
	models.InitDB()
	setRouter()
}

func setRouter() {

	router := gin.Default()

	catV1 := router.Group("api/v1/category")
	{
		catV1.POST("/", services.SaveCategory)
		catV1.GET("/", services.GetAllCategory)
		catV1.GET("/:id", services.GetOneCategory)
		catV1.PUT("/:id", services.UpdateCategory)
	}

	goodV1 := router.Group("/api/v1/good")
	{
		goodV1.POST("/", services.SaveGood)
		goodV1.GET("/", services.GetAllGood)
		goodV1.GET("/:id", services.GetOneGood)
		goodV1.PUT("/:id", services.UpdateGood)
	}

	router.Run()
}
