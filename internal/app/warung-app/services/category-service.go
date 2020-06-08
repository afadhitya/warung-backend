package services

import (
	"net/http"
	"strconv"

	"github.com/afadhitya/warung-app/internal/app/warung-app/models"
	"github.com/gin-gonic/gin"
)

func SaveCategory(c *gin.Context) {

	category := getAttribute(c)

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

func GetOneCategory(c *gin.Context) {
	id, _ := strconv.ParseInt((c.Param("id")), 10, 64)
	category := getOneCategory(int(id))

	if isNotFound(&category, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   category,
	})
}

func UpdateCategory(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	category := getOneCategory(int(id))

	if isNotFound(&category, c) {
		return
	}

	categoryUpdated := getAttribute(c)
	models.DB.Model(&category).Updates(categoryUpdated)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   category,
	})
}

func getOneCategory(id int) models.Category {
	var category models.Category
	models.DB.Debug().Preload("Goods").First(&category, id)
	return category
}

func getAttribute(c *gin.Context) models.Category {
	category := models.Category{
		Name:       c.PostForm("name"),
		CodeNumber: c.PostForm("codeNumber"),
	}

	return category
}

func isNotFound(cat *models.Category, c *gin.Context) bool {
	if cat.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Category not found",
		})
	}

	return (cat.ID == 0)
}
