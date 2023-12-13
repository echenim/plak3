package repositories

import (
	"database/sql"

	"github.com/plak3com/plak3/internal/models/views"
)

type Plak3UserEmailRepository struct {
	db *sql.DB
}

func NewPlak3UserEmailRepository(_db *sql.DB) *Plak3UserEmailRepository {
	return &Plak3UserEmailRepository{db: _db}
}

func (r *Plak3UserEmailRepository) StartOnboarding(model views.Plak3UserEmail) (bool, error) {
	// SQL statement to insert a new user
	query := `INSERT INTO plak_users_email (email,email_confirmed) VALUES ($1, $2) RETURNING id`

	// Use db.QueryRow to execute the query
	err := r.db.QueryRow(query, model.Email, false).Scan(&model.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}
