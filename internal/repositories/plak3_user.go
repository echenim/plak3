package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/plak3com/plak3/internal/models/entities"
	"github.com/plak3com/plak3/internal/models/searchmodels"
)

type Plak3UserRepository struct {
	db *sql.DB
}

func NewPlak3UserRepository() *Plak3UserRepository {
	return &Plak3UserRepository{}
}

func (r *Plak3UserRepository) Get(ctx context.Context) ([]entities.PlakUser, error) {
	return []entities.PlakUser{}, nil
}

func (r *Plak3UserRepository) Find(ctx context.Context) (entities.PlakUser, error) {
	return entities.PlakUser{}, nil
}

func (r *Plak3UserRepository) Search(param searchmodels.UserSearchCriteria) ([]entities.PlakUser, error) {
	fmt.Println(param)
	return []entities.PlakUser{}, nil
}

func (r *Plak3UserRepository) Save(ctx context.Context) (entities.PlakUser, error) {
	return entities.PlakUser{}, nil
}

func (r *Plak3UserRepository) Edit(ctx context.Context) (entities.PlakUser, error) {
	return entities.PlakUser{}, nil
}

func (r *Plak3UserRepository) Remove(ctx context.Context) error {
	return nil
}
