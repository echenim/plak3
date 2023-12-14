package services

import (
	"github.com/plak3com/plak3/internal/models/views"
	"github.com/plak3com/plak3/internal/repositories"
)

type Plak3EmailBankService struct {
	repo *repositories.Plak3EmailBankRepository
}

func NewPlak3EmailBankService(_repo *repositories.Plak3UserRepository) *Plak3UserService {
	return &Plak3UserService{repo: _repo}
}

func (s *Plak3EmailBankService) Save(u views.Plak3EmailBank) (bool, error) {
	return s.repo.StartOnboarding(u)
}
