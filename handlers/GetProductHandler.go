package handlers

import (
	"github.com/gin-gonic/gin"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"
)

func GetProduct(c *gin.Context) {

	productId := c.Param("id")

	var product entities.Product
	database.Instance.First(&product, productId)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
