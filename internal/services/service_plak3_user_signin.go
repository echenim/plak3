package services

import (
	"context"

	"github.com/plak3com/plak3/internal/repositories"
)

type Plak3UserSignInService struct {
	repo repositories.Plak3UserSignInRepository
}

func NewPlak3UserSignInService(_repo repositories.Plak3UserSignInRepository) *Plak3UserSignInService {
	return &Plak3UserSignInService{repo: _repo}
}

func (s *Plak3UserSignInService) SignIn(ctx context.Context) error {
	return s.repo.SignIn(ctx)
}

func (s *Plak3UserSignInService) Revoke(ctx context.Context) error {
	return s.repo.Revoke(ctx)
}

func (s *Plak3UserSignInService) Remove(ctx context.Context) error {
	return s.repo.Remove(ctx)
}

func (s *Plak3UserSignInService) CreateSignIn(ctx context.Context) error {
	return s.repo.CreateSignIn(ctx)
}
