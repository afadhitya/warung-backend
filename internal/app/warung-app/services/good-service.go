package services

import (
	"net/http"
	"strconv"

	"github.com/afadhitya/warung-app/internal/app/warung-app/models"
	"github.com/gin-gonic/gin"
)

func SaveGood(c *gin.Context) {
	name := c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))
	catID, _ := strconv.ParseUint((c.PostForm("catID")), 10, 64)
	stock, _ := strconv.ParseInt((c.PostForm("stock")), 10, 64)
	discount, _ := strconv.ParseFloat((c.PostForm("discount")), 64)

	good := models.Good{
		Name:       name,
		Price:      price,
		Stock:      stock,
		Discount:   discount,
		CategoryID: uint(catID),
	}

	models.DB.Save(&good)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Category has saved to DB",
		"resourceId": good.ID,
	})
}

func GetAllGood(c *gin.Context) {
	// var category models.Category
	var goods []models.Good
	models.DB.Debug().Preload("Category").Find(&goods)

	// .Related(&category)
	if len(goods) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No good found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   goods,
	})
}

func GetOneGood(c *gin.Context) {
	id, _ := strconv.ParseInt((c.Param("id")), 10, 64)
	good := getOne(id)

	if good.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No good found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   good,
	})
}

func getOne(id int64) models.Good {
	var good models.Good
	models.DB.Debug().Preload("Category").First(&good, id)
	return good
}
