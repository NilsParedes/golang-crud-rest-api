package handlers

import (
	"github.com/gin-gonic/gin"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"net/http"
)

func GetProducts(c *gin.Context) {

	var products []entities.Product
	database.Instance.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}
