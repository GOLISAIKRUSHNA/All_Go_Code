package repositorys

import (
	 "gorm.io/gorm"
	 "collection_api/models"
)


type LoanRepository interface{
	
    Create(loan *models.Loan) error
    Update(loan *models.Loan) error
    // FindByID(id uint) (*models.User, error)
    // Delete(id uint) error
}

type loanRepository struct {
    db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
    return &loanRepository{db: db}
}


func (r *loanRepository) Create(user *models.Loan) error {
    return r.db.Create(user).Error
}


func (r *loanRepository) Update(user *models.Loan) error {
    return r.db.Save(user).Error
}