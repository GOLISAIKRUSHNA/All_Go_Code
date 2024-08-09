package handlers

import (
	"net/http"

	dto "collection/dto"


	"github.com/gin-gonic/gin"
)

func CreateLoan(c *gin.Context) {

	var loanRequest dto.CreateLoanRequest

	// Bind the JSON body to the LoanRequest struct
	if err := c.ShouldBindJSON(&loanRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// Send a JSON response with the loan details
	c.JSON(http.StatusOK, loanRequest)
}
