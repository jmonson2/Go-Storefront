package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"storefront/db"
	"storefront/items"
)

func GetAllInventory(context *gin.Context) {
	inventory, err := db.GetAllInventory()
	if err != nil {
		context.JSON(500, err)
	}
	context.JSON(200, inventory)
}

func GetInventoryBySKU(context *gin.Context) {
	inventory, err := db.GetInventoryBySKU(context.Query("sku"))
	if err != nil {
		context.JSON(500, gin.H{
			"Error":   "Internal server error",
			"Message": "Error retrieving inventory" + err.Error(),
		})
	}
	context.JSON(200, inventory)
}

func AddInventory(context *gin.Context) {
	var inventory items.Inventory
	err := context.BindJSON(&inventory)
	if err != nil {
		fmt.Println(err)
		context.JSON(400, gin.H{
			"Error":   "Bad request",
			"Message": "Invalid JSON" + err.Error(),
		})
		return
	}
	err = db.InsertInventory(inventory)
	if err != nil {
		context.JSON(500, gin.H{
			"Error":   "Internal server error",
			"Message": "Error inserting inventory: " + err.Error(),
		})
	}
	context.JSON(200, inventory)
}
