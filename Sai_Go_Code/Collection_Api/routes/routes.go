package routes

import (

	"collection/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Initialize a new Gin router
	router := gin.Default()

	// Define a route for loan-related operations
	//router.GET("/loan", handlers.GetLoanDetails)
	router.POST("/loan", handlers.CreateLoan)
	// You can add more routes here as needed

	return router
}
