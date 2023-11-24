package repositories

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/plak3com/plak3/internal/models/searchmodels"
	"github.com/plak3com/plak3/internal/models/views"
)

type Plak3UserRepository struct {
	db *sql.DB
}

func NewPlak3UserRepository(_db *sql.DB) *Plak3UserRepository {
	return &Plak3UserRepository{db: _db}
}

func (r *Plak3UserRepository) Get() ([]views.PlakUser, error) {
	var users []views.PlakUser

	// SQL query to select all users
	query := "SELECT id, first_name,last_name, email FROM plak_users"

	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan each row into a PlakUser struct
	for rows.Next() {
		var user views.PlakUser
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Plak3UserRepository) Find(idparam int64) (views.PlakUser, error) {
	var user views.PlakUser

	// SQL query to select the user by ID
	query := "SELECT id, first_name,last_name, email FROM plak_users WHERE id = $1"

	// Execute the query
	row := r.db.QueryRow(query, idparam)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// Optional: return a custom error indicating that no user was found
			return views.PlakUser{}, err
		}
		// Return the error encountered during query execution or scanning
		return views.PlakUser{}, err
	}

	return user, nil
}

func (r *Plak3UserRepository) Search(param searchmodels.UserSearchCriteria) ([]views.PlakUser, bool, error) {
	var results []views.PlakUser

	// Using a string builder for efficient query construction
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT id, first_name, last_name, email FROM plak_users")

	args := make([]interface{}, 0)
	conditions := make([]string, 0)

	// Dynamically building the query based on search criteria
	if param.FirstName != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(conditions)+1))
		args = append(args, param.FirstName)
	}
	if param.Email != "" {
		conditions = append(conditions, fmt.Sprintf("email = $%d", len(conditions)+1))
		args = append(args, param.Email)
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}

	// Executing the query
	rows, err := r.db.Query(queryBuilder.String(), args...)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	// Scanning rows into the results
	for rows.Next() {
		var user views.PlakUser
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, false, err
		}
		results = append(results, user)
	}

	if err := rows.Err(); err != nil {
		return nil, false, err
	}

	return results, len(results) > 0, nil
}

func (r *Plak3UserRepository) Edit(u views.PlakUser) (views.PlakUser, error) {
	// SQL statement to update the user's details
	query := "UPDATE plak_users SET first_name = $1,last_name=$2, email = $3, lastupdated = $4, id = $5"

	// Execute the query
	_, err := r.db.Exec(query, u.FirstName, u.LastName, u.Email, time.Now(), u.ID)
	if err != nil {
		return views.PlakUser{}, err
	}

	return u, nil
}

func (r *Plak3UserRepository) Remove(idParam int64) error {
	// SQL statement to delete the user
	query := "DELETE FROM plak_users WHERE id = $1"

	// Execute the query
	_, err := r.db.Exec(query, idParam)
	return err
}

func (r *Plak3UserRepository) Save(u views.PlakUser) (views.PlakUser, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return u, err
	}

	// Save user
	nu, err := saveUser(tx, u)
	if err != nil {
		tx.Rollback()
		return nu, err
	}

	// Save roles
	if err := addUserToRoles(tx, nu); err != nil {
		tx.Rollback()
		return u, err
	}

	return u, tx.Commit()
}
