package services

import(
	"collection_api/models"
	"collection_api/repositorys"
	"collection_api/dto"
	
)

type LoanService interface {
    CreateUser(req dto.CreateLoanRequest) (*models.Loan, error)
    // GetUser(id uint) (*models.User, error)
    UpdateUser(id uint, req dto.UpdateLoanRequest) (*models.Loan, error)
    // DeleteUser(id uint) error
}

type loanService struct {
    loanRepo repositorys.LoanRepository
}

func NewUserService(userRepo repositorys.LoanRepository) *loanService {
    return &loanService{loanRepo: userRepo}
}

func (s *loanService) CreateUser(req dto.CreateLoanRequest) (*models.Loan, error) {
    loan := &models.Loan{









        // Name:  req.Name,
        // Email: req.Email,
        // Age:   req.Age,
    }
    if err := s.loanRepo.Create(user); err != nil {
        return nil, err
    }
    return user, nil
}

