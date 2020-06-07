package services

import (
	"net/http"

	"github.com/afadhitya/warung-app/internal/app/warung-app/models"
	"github.com/gin-gonic/gin"
)

func SaveCategory(c *gin.Context) {
	category := models.Category{
		Name:       c.PostForm("name"),
		CodeNumber: c.PostForm("codeNumber"),
	}

	models.DB.Save(&category)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Category has saved to DB",
		"resourceId": category.ID,
	})
}

func GetAllCategory(c *gin.Context) {
	var categories []models.Category
	// var goods []models.Good
	models.DB.Debug().Preload("Goods").Find(&categories)

	if len(categories) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No category found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   categories,
	})
}
