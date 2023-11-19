package repositories

import "context"

type Plak3UserSignInRepository struct{}

func NewPlak3UserSignInRepository() *Plak3UserSignInRepository {
	return &Plak3UserSignInRepository{}
}

func (r *Plak3UserSignInRepository) SignIn(ctx context.Context) error {
	return nil
}

func (r *Plak3UserSignInRepository) Revoke(ctx context.Context) error {
	return nil
}

func (r *Plak3UserSignInRepository) Remove(ctx context.Context) error {
	return nil
}

func (r *Plak3UserSignInRepository) CreateSignIn(ctx context.Context) error {
	return nil
}
