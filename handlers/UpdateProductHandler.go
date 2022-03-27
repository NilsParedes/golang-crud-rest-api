package handlers

import (
	"github.com/gin-gonic/gin"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"golang-crud-rest-api/requests"
	"net/http"
)

func UpdateProduct(c *gin.Context) {

	productId := c.Param("id")
	var request requests.UpdateProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product entities.Product

	database.Instance.First(&product, productId)
	database.Instance.Model(&product).Updates(entities.Product{Name: request.Name, Price: request.Price, Description: request.Description})

	c.JSON(http.StatusOK, gin.H{"data": product})
}
