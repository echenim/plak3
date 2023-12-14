package repositories

import (
	"database/sql"
	"time"

	"github.com/plak3com/plak3/internal/models/views"
)

type Plak3EmailBankRepository struct {
	db *sql.DB
}

func NewPlak3EmailBankRepository(_db *sql.DB) *Plak3EmailBankRepository {
	return &Plak3EmailBankRepository{db: _db}
}

func (r *Plak3EmailBankRepository) StartOnboarding(m views.Plak3EmailBank) (bool, error) {
	// SQL statement to insert a new user
	query := `INSERT INTO plak_email_bank (email) VALUES ($1) RETURNING id`

	// Use db.QueryRow to execute the query
	err := r.db.QueryRow(query, m.Email, false).Scan(&m.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Plak3EmailBankRepository) ValidateEmail(m string) (views.Plak3EmailBank, error) {
	// SQL statement to insert a new user
	query := `SELECT id,email plak_email_bank WHERE  id== $1`
	u := views.Plak3EmailBank{}

	// Execute the query
	row := r.db.QueryRow(query, m)
	err := row.Scan(&u.Id, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// Optional: return a custom error indicating that no user was found
			return u, err
		}
		// Return the error encountered during query execution or scanning
		return u, err
	}

	err = r.Updatevalidation(u.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			// Optional: return a custom error indicating that no user was found
			return u, err
		}
		// Return the error encountered during query execution or scanning
		return u, err
	}

	return u, nil
}

func (r *Plak3EmailBankRepository) Updatevalidation(u string) error {
	// SQL statement to update the user's details
	query := `UPDATE plak_email_bank  SET lastUpdated = $1 WHERE id=$2`

	// Execute the query
	_, err := r.db.Exec(query, time.Now(), u)
	if err != nil {
		return err
	}

	return nil
}
