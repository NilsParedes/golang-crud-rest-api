package main

import (
	"github.com/gin-gonic/gin"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/routes"
)

func main() {

	// Init database
	database.Init()

	// Init
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start server
	_ = router.Run()

}
