package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/plak3com/plak3/internal/models/entities"
	"github.com/plak3com/plak3/internal/models/searchmodels"
)

type Plak3UserTokensRepository struct {
	db *sql.DB
}

func NewPlak3UserTokensRepository() *Plak3UserTokensRepository {
	return &Plak3UserTokensRepository{}
}

func (r *Plak3UserTokensRepository) Get(ctx context.Context) ([]entities.Plak3UserTokens, error) {
	return []entities.Plak3UserTokens{}, nil
}

func (r *Plak3UserTokensRepository) Find(ctx context.Context) (entities.Plak3UserTokens, error) {
	return entities.Plak3UserTokens{}, nil
}

func (r *Plak3UserTokensRepository) Search(param searchmodels.UserSearchCriteria) ([]entities.Plak3UserTokens, error) {
	fmt.Println(param)
	return []entities.Plak3UserTokens{}, nil
}

func (r *Plak3UserTokensRepository) Save(ctx context.Context) (entities.Plak3UserTokens, error) {
	return entities.Plak3UserTokens{}, nil
}

func (r *Plak3UserTokensRepository) Edit(ctx context.Context) (entities.Plak3UserTokens, error) {
	return entities.Plak3UserTokens{}, nil
}

func (r *Plak3UserTokensRepository) Remove(ctx context.Context) error {
	return nil
}
