package services

import (
	"github.com/afadhitya/warung-backend/internal/app/warung-app/config"
	"github.com/afadhitya/warung-backend/internal/app/warung-app/util"
	"net/http"
	"strconv"

	"github.com/afadhitya/warung-backend/internal/app/warung-app/models"
	"github.com/gin-gonic/gin"
)

func SaveGood(c *gin.Context) {
	good := getGoodAttributes(c)

	config.DB.Save(&good)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Category has saved to DB",
		"resourceId": good.ID,
	})
}

func GetAllGood(c *gin.Context) {
	var goods []models.Good
	config.DB.Debug().Preload("Category").Find(&goods)

	if len(goods) <= 0 {
		util.SetContextError(c, http.StatusNotFound, "No good found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   goods,
	})
}

func GetOneGood(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	good := getOneGood(id)

	if isGoodNotFound(&good, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   good,
	})
}

func UpdateGood(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	good := getOneGood(id)

	if isGoodNotFound(&good, c) {
		return
	}

	goodUpdated := getGoodAttributes(c)
	config.DB.Model(&good).Updates(goodUpdated)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   good,
	})
}

func DeleteGood(c *gin.Context){
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	good := getOneGood(id)

	if isGoodNotFound(&good, c) {
		return
	}

	config.DB.Debug().Delete(&good)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   good,
	})
}

func getOneGood(id int64) models.Good {
	var good models.Good
	config.DB.Debug().Preload("Category").First(&good, id)
	return good
}

func isGoodNotFound(good *models.Good, c *gin.Context) bool {
	if good.ID == 0 {
		util.SetContextError(c, http.StatusNotFound, "Item not found")
	}
	return good.ID == 0
}

func getGoodAttributes(c *gin.Context) models.Good {
	name := c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))
	catID, _ := strconv.ParseUint(c.PostForm("catID"), 10, 64)
	stock, _ := strconv.ParseInt(c.PostForm("stock"), 10, 64)
	discount, _ := strconv.ParseFloat(c.PostForm("discount"), 64)

	good := models.Good{
		Name:       name,
		Price:      price,
		Stock:      stock,
		Discount:   discount,
		CategoryID: uint(catID),
	}

	return good
}
