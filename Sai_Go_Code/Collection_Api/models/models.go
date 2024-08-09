package models

import (
	"time"

	"gorm.io/gorm" //it is a library for db
)

type Loan struct {
	gorm.Model
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
	// Name                string `json:"name"`
	// Email               string `json:"email" gorm:"unique"`
	// Age                 int    `json:"age"`
}
