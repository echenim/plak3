package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/plak3com/plak3/internal/models/entities"
	"github.com/plak3com/plak3/internal/models/searchmodels"
)

type Plak3UserLoginsRepository struct {
	db *sql.DB
}

func NewPlak3UserLoginsRepository() *Plak3UserLoginsRepository {
	return &Plak3UserLoginsRepository{}
}

func (r *Plak3UserLoginsRepository) Get(ctx context.Context) ([]entities.Plak3UserLogins, error) {
	return []entities.Plak3UserLogins{}, nil
}

func (r *Plak3UserLoginsRepository) Find(ctx context.Context) (entities.Plak3UserLogins, error) {
	return entities.Plak3UserLogins{}, nil
}

func (r *Plak3UserLoginsRepository) Search(param searchmodels.UserSearchCriteria) ([]entities.Plak3UserLogins, error) {
	fmt.Println(param)
	return []entities.Plak3UserLogins{}, nil
}

func (r *Plak3UserLoginsRepository) Save(ctx context.Context) (entities.Plak3UserLogins, error) {
	return entities.Plak3UserLogins{}, nil
}

func (r *Plak3UserLoginsRepository) Edit(ctx context.Context) (entities.Plak3UserLogins, error) {
	return entities.Plak3UserLogins{}, nil
}

func (r *Plak3UserLoginsRepository) Remove(ctx context.Context) error {
	return nil
}
