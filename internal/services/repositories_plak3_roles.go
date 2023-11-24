package services

import (
	"github.com/plak3com/plak3/internal/models/views"
	"github.com/plak3com/plak3/internal/repositories"
)

type Plak3RolesServices struct {
	repo *repositories.Plak3RolesRepository
}

func NewPPlak3RolesServices(_repo *repositories.Plak3RolesRepository) *Plak3RolesServices {
	return &Plak3RolesServices{repo: _repo}
}

func (r *Plak3RolesServices) Get() ([]views.Plak3Roles, error) {
	return r.repo.Get()
}

func (r *Plak3RolesServices) Find(id string) (views.Plak3Roles, error) {
	return r.repo.Find(id)
}

func (r *Plak3RolesServices) Search(role_id []string) ([]views.Plak3Roles, error) {
	return r.repo.Search(role_id)
}
