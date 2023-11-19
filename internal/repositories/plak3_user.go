package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/plak3com/plak3/internal/models"
	"github.com/plak3com/plak3/internal/models/searchmodels"
)

type Plak3UserRepository struct {
	db *sql.DB
}

func NewPlak3UserRepository() *Plak3UserRepository {
	return &Plak3UserRepository{}
}

func (r *Plak3UserRepository) Get(ctx context.Context) ([]models.PlakUser, error) {
	return []models.PlakUser{}, nil
}

func (r *Plak3UserRepository) Find(ctx context.Context) (models.PlakUser, error) {
	return models.PlakUser{}, nil
}

func (r *Plak3UserRepository) Search(param searchmodels.UserSearchCriteria) ([]models.PlakUser, error) {
	fmt.Println(param)
	return []models.PlakUser{}, nil
}

func (r *Plak3UserRepository) Save(ctx context.Context) (models.PlakUser, error) {
	return models.PlakUser{}, nil
}

func (r *Plak3UserRepository) Edit(ctx context.Context) (models.PlakUser, error) {
	return models.PlakUser{}, nil
}

func (r *Plak3UserRepository) Remove(ctx context.Context) error {
	return nil
}
