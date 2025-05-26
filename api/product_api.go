package api

import (
	"github.com/gin-gonic/gin"
	"storefront/db"
	"storefront/items"
	"strconv"
)

func GetAllProducts(context *gin.Context) {
	products, err := db.GetAllProducts()
	if err != nil {
		context.JSON(500, gin.H{
			"Error":   "Internal server error",
			"Message": "Error retrieving products" + err.Error(),
		})
	}
	context.JSON(200, products)
}

func GetProductsBySKU(context *gin.Context) {
	sku := context.Query("sku")

	products, err := db.GetProductBySKU(sku)
	if err != nil {
		context.JSON(500, gin.H{
			"Error":   "Internal server error",
			"Message": "Error retrieving products" + err.Error(),
		})
		return
	}
	context.JSON(200, products)

}

func GetProductByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		context.JSON(400, gin.H{
			"Error":   "Bad request",
			"Message": "ID is required" + err.Error(),
		})
		return
	}
	products, err := db.GetProductByID(id)
	if err != nil {
		context.JSON(500, gin.H{
			"Error":   "Internal server error",
			"Message": "Error retrieving products" + err.Error(),
		})
		return
	}
	context.JSON(200, products)
}

func GetProductsByName(context *gin.Context) {
	name := context.Query("name")
	products, err := db.GetProductsByName(name)
	if err != nil {
		context.JSON(500, gin.H{
			"Error":   "Internal server error",
			"Message": "Error retrieving products" + err.Error(),
		})
		return
	}
	context.JSON(200, products)
}

func AddProduct(context *gin.Context) {
	var product items.Product
	err := context.BindJSON(&product)
	if err != nil {
		context.JSON(400, gin.H{
			"Error":   "Bad request",
			"Message": "Invalid JSON" + err.Error(),
		})
		return
	}

	err = db.InsertProduct(product)
	if err != nil {
		context.JSON(500, gin.H{
			"Error":   "Internal server error",
			"Message": "Error inserting product" + err.Error(),
		})
	}
	context.JSON(200, product)
}

func SellProduct(context *gin.Context) {
	var product items.Product
	err := context.BindJSON(&product)
	if err != nil {
		context.JSON(400, gin.H{
			"Error":   "Bad request",
			"Message": "Invalid JSON" + err.Error(),
		})
		return
	}

	err = db.SellProduct(product)
	if err != nil {
		context.JSON(500, gin.H{
			"Error":   "Internal server error",
			"Message": "Error selling product" + err.Error(),
		})
	}
	context.JSON(200, product)
}
