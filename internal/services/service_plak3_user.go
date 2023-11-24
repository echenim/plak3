package services

import (
	"github.com/plak3com/plak3/internal/models/searchmodels"
	"github.com/plak3com/plak3/internal/models/views"
	"github.com/plak3com/plak3/internal/repositories"
)

type Plak3UserService struct {
	repo *repositories.Plak3UserRepository
}

func NewPlak3UserService(_repo *repositories.Plak3UserRepository) *Plak3UserService {
	return &Plak3UserService{repo: _repo}
}

func (s *Plak3UserService) Get() ([]views.PlakUser, error) {
	return s.repo.Get()
}

func (s *Plak3UserService) Find(idParam int64) (views.PlakUser, error) {
	return s.repo.Find(idParam)
}

func (s *Plak3UserService) Search(param searchmodels.UserSearchCriteria) ([]views.PlakUser, bool, error) {
	return s.repo.Search(param)
}

func (s *Plak3UserService) Save(u views.PlakUser) (views.PlakUser, error) {
	return s.repo.Save(u)
}

func (s *Plak3UserService) Edit(u views.PlakUser) (views.PlakUser, error) {
	return s.repo.Edit(u)
}

func (r *Plak3UserService) Remove(idParam int64) error {
	return r.repo.Remove(idParam)
}
