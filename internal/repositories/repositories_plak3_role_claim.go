package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/plak3com/plak3/internal/models/entities"
	"github.com/plak3com/plak3/internal/models/searchmodels"
)

type Plak3RoleClaimsRepository struct {
	db *sql.DB
}

func NewPlak3RoleClaimsRepository() *Plak3RoleClaimsRepository {
	return &Plak3RoleClaimsRepository{}
}

func (r *Plak3RoleClaimsRepository) Get(ctx context.Context) ([]entities.Plak3RoleClaims, error) {
	return []entities.Plak3RoleClaims{}, nil
}

func (r *Plak3RoleClaimsRepository) Find(ctx context.Context) (entities.Plak3RoleClaims, error) {
	return entities.Plak3RoleClaims{}, nil
}

func (r *Plak3RoleClaimsRepository) Search(param searchmodels.UserSearchCriteria) ([]entities.Plak3RoleClaims, error) {
	fmt.Println(param)
	return []entities.Plak3RoleClaims{}, nil
}

func (r *Plak3RoleClaimsRepository) Save(ctx context.Context) (entities.Plak3RoleClaims, error) {
	return entities.Plak3RoleClaims{}, nil
}

func (r *Plak3RoleClaimsRepository) Edit(ctx context.Context) (entities.Plak3RoleClaims, error) {
	return entities.Plak3RoleClaims{}, nil
}

func (r *Plak3RoleClaimsRepository) Remove(ctx context.Context) error {
	return nil
}
