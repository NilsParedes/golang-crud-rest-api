package handlers

import (
	"github.com/gin-gonic/gin"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/entities"
	"golang-crud-rest-api/requests"
	"net/http"
)

func CreateProduct(c *gin.Context) {

	var request requests.CreateProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := entities.Product{Name: request.Name, Price: request.Price, Description: request.Description}
	database.Instance.Create(&product)

	c.JSON(http.StatusCreated, gin.H{"data": product})
}
