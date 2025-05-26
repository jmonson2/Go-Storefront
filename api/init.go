package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	router := gin.Default()

	initProduct(router)
	initInventory(router)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Println("Error initializing router: ", err)
	}
}

func initProduct(router *gin.Engine) {
	router.GET("/products", GetAllProducts)
	router.GET("/products/sku", GetProductsBySKU)
	router.GET("/products/id", GetProductByID)
	router.GET("/products/name", GetProductsByName)
	router.POST("/products/add", AddProduct)
	router.POST("/products/sell", SellProduct)
}

func initInventory(router *gin.Engine) {
	router.GET("/inventory", GetAllInventory)
	router.GET("/inventory/sku", GetInventoryBySKU)
	router.POST("/inventory/add", AddInventory)
}
