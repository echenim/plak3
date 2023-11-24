package repositories

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/plak3com/plak3/internal/models/views"
)

func saveUser(tx *sql.Tx, u views.PlakUser) (views.PlakUser, error) {
	// SQL statement to insert a new user
	query := `INSERT INTO plak_users (first_name,last_name,email) VALUES ($1, $2,$3) RETURNING id`

	// Use db.QueryRow to execute the query
	err := tx.QueryRow(query, u.FirstName, u.LastName, u.Email).Scan(&u.ID)
	if err != nil {
		return views.PlakUser{}, err
	}

	return u, nil
}

func addUserToRoles(tx *sql.Tx, u views.PlakUser) error {
	err := removeUserRoles(tx, u.ID)
	if err != nil {
		return err
	}
	return saveUserRoles(tx, u.ID, u.Roles)
}

func removeUserRoles(tx *sql.Tx, u int64) error {
	// SQL statement to delete the user
	query := "DELETE FROM plak_user_roles WHERE user_id = $1"

	// Execute the query
	_, err := tx.Exec(query, u)
	return err
}

func saveUserRoles(tx *sql.Tx, user_id int64, u []string) error {
	// SQL statement to insert a new user
	query := `INSERT INTO  plak_user_roles (user_id,role_id) VALUES ($1, $2)`
	var errStrs []string
	for _, i := range u {
		// Use db.Exec to execute the query without returning the id
		_, err := tx.Exec(query, user_id, i)
		if err != nil {
			errStrs = append(errStrs, err.Error())
		}
	}

	if len(errStrs) > 0 {
		return errors.New(strings.Join(errStrs, ", "))
	}
	return nil
}
