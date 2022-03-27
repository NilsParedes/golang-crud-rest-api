package routes

import (
	"github.com/gin-gonic/gin"
	"golang-crud-rest-api/handlers"
)

func RegisterRoutes(router *gin.Engine) {

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("/", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
			products.POST("/", handlers.CreateProduct)
			products.PUT("/:id", handlers.UpdateProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
		}
	}

}
