package services

import (
	"github.com/afadhitya/warung-backend/internal/app/warung-app/config"
	"github.com/afadhitya/warung-backend/internal/app/warung-app/models"
)

// func DeleteItemFromOrder(c *gin.Context) {

// 	goodId, _ := strconv.ParseInt(c.Param("goodId"), 10, 64)
// 	itemId, _ := strconv.ParseInt(c.Param("itemId"), 10, 64)

// 	order := getItemByOrderIDAndItemID()
// }

func getItemByID(id int) models.Item {
	var item models.Item
	config.DB.Debug().First(&item, id)
	return item
}

func saveItem(item *models.Item) {
	if itemAvailable(item.OrderID, item.GoodID) {
		var itemFromDB = getItemByOrderIDAndItemID(item.OrderID, item.GoodID)
		itemFromDB.Amount = item.Amount + itemFromDB.Amount
		config.DB.Save(&itemFromDB)
		return
	}
	config.DB.Create(item)
}

func itemAvailable(orderID uint, goodID uint) bool {
	var count int
	var item models.Item
	config.DB.Debug().Where(&models.Item{
		OrderID: uint(orderID),
		GoodID:  uint(goodID),
	}).Find(&item).Count(&count)

	return count > 0
}

func getItemByOrderIDAndItemID(orderID uint, goodID uint) models.Item {
	var item models.Item
	config.DB.Debug().Where(&models.Item{
		OrderID: uint(orderID),
		GoodID:  uint(goodID),
	}).First(&item)
	return item
}

// func deleteItemByOrderIDAndItemID(orderID uint, goodID uint) {

// }
