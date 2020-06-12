package main

import (
	"fmt"
	"github.com/afadhitya/warung-backend/internal/app/warung-app/config"

	"github.com/afadhitya/warung-backend/internal/app/warung-app/services"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("application start")
	config.InitDB()
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
		goodV1.DELETE("/:id", services.DeleteGood)
	}

	userV1 := router.Group("/api/v1/user")
	{
		userV1.POST("/signup", services.SignUp)
		userV1.POST("/login", services.SignIn)
	}

	router.Run()
}
