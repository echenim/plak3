package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/plak3com/plak3/internal/models/entities"
	"github.com/plak3com/plak3/internal/models/searchmodels"
)

type Plak3UserClaimsRepository struct {
	db *sql.DB
}

func NewPlak3UserClaimsRepository() *Plak3UserClaimsRepository {
	return &Plak3UserClaimsRepository{}
}

func (r *Plak3UserClaimsRepository) Get(ctx context.Context) ([]entities.Plak3UserClaims, error) {
	return []entities.Plak3UserClaims{}, nil
}

func (r *Plak3UserClaimsRepository) Find(ctx context.Context) (entities.Plak3UserClaims, error) {
	return entities.Plak3UserClaims{}, nil
}

func (r *Plak3UserClaimsRepository) Search(param searchmodels.UserSearchCriteria) ([]entities.Plak3UserClaims, error) {
	fmt.Println(param)
	return []entities.Plak3UserClaims{}, nil
}

func (r *Plak3UserClaimsRepository) Save(ctx context.Context) (entities.Plak3UserClaims, error) {
	return entities.Plak3UserClaims{}, nil
}

func (r *Plak3UserClaimsRepository) Edit(ctx context.Context) (entities.Plak3UserClaims, error) {
	return entities.Plak3UserClaims{}, nil
}

func (r *Plak3UserClaimsRepository) Remove(ctx context.Context) error {
	return nil
}
