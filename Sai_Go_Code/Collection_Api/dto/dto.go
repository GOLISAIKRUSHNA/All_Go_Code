package dto

import "time"

type CreateLoanRequest struct {
	// Id                  int
	Loan_ref_number     string   `json:"loan_ref_number"`
	Customer_ref_number string   `json:"customer_ref_number"`
	Loan_ref_numberoan_amount  float64   `json:"loan_amount"` 
	Enum                string    ` json:"enum"`
	Principal           float64   `json:"principal"`
	Interest            float64  `json:"interest"` 
	Fees_due            float64   `json:"fees_due"`
	Is_deleted          int      `json:"is_deleted"`
	Created_at          time.Time   `json:"created_at"`
	Updated_at          time.Time `json:"updated_at"`

	// Name  string `json:"name" binding:"required"`
	// Email string `json:"email" binding:"required,email"`
	// Age   int    `json:"age" binding:"required"`
}

type UpdateLoanRequest struct {
	
	Id                  int
	loan_ref_number     string
	customer_ref_number string
	loan_amount         float64
	enum                string
	principal           float64
	interest            float64
	fees_due            float64
	is_deleted          int
	created_at          time.Time
	updated_at          time.Time
	
    // Name  string `json:"name"`
    // Email string `json:"email" binding:"email"`
    // Age   int    `json:"age"`
}
