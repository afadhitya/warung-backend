package services

import (
	"net/http"
	"strconv"

	"github.com/afadhitya/warung-backend/internal/app/warung-app/config"
	"github.com/afadhitya/warung-backend/internal/app/warung-app/models"
	"github.com/gin-gonic/gin"
)

func AddItem(c *gin.Context) {
	// orderId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	goodId, _ := strconv.ParseInt(c.Param("goodId"), 10, 64)
	userId, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	amount, _ := strconv.ParseInt(c.PostForm("amount"), 10, 64)

	order := getActiveCart(int(userId))

	if order.ID == 0 {
		saveNewOrder(&order, userId)
	}
	good := getOneGood(goodId)
	item := models.Item{
		Order:   order,
		OrderID: order.ID,
		Good:    good,
		GoodID:  good.ID,
		Amount:  int(amount),
	}

	saveItem(&item)
	orderNew := getOneOrder(int64(order.ID))
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   orderNew,
	})
}

func getOneOrder(id int64) models.Order {
	var order models.Order
	config.DB.Debug().Preload("User").Preload("Items").First(&order, id)
	return order
}

func saveNewOrder(order *models.Order, userID int64) {
	order.Status = models.STATUS_ON_CART
	order.UserID = int(userID)
	config.DB.Debug().Create(order)
}

func getActiveCart(userID int) models.Order {
	var order models.Order
	config.DB.Where(&models.Order{
		UserID: userID,
		Status: models.STATUS_ON_CART,
	}).First(&order)
	return order
}
