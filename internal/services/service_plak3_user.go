package services

import (
	"context"

	"github.com/plak3com/plak3/internal/models"
	"github.com/plak3com/plak3/internal/models/searchmodels"
	"github.com/plak3com/plak3/internal/repositories"
)

type Plak3UserService struct {
	repo repositories.Plak3UserRepository
}

func NewPlak3UserService(_repo repositories.Plak3UserRepository) *Plak3UserService {
	return &Plak3UserService{repo: _repo}
}

func (s *Plak3UserService) Get(ctx context.Context) ([]models.PlakUser, error) {
	return s.repo.Get(ctx)
}

func (s *Plak3UserService) FindById(ctx context.Context) (models.PlakUser, error) {
	return s.repo.Find(ctx)
}

func (s *Plak3UserService) Search(param searchmodels.UserSearchCriteria) ([]models.PlakUser, error) {
	return s.repo.Search(param)
}

func (s *Plak3UserService) Save(ctx context.Context) (models.PlakUser, error) {
	return s.repo.Save(ctx)
}

func (s *Plak3UserService) Edit(ctx context.Context) (models.PlakUser, error) {
	return s.repo.Edit(ctx)
}

func (r *Plak3UserService) Remove(ctx context.Context) error {
	return r.repo.Remove(ctx)
}
