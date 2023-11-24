package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/plak3com/plak3/internal/models/views"
)

type Plak3RolesRepository struct {
	db *sql.DB
}

func NewPlak3RolesRepository() *Plak3RolesRepository {
	return &Plak3RolesRepository{}
}

func (r *Plak3RolesRepository) Get() ([]views.Plak3Roles, error) {
	var roles []views.Plak3Roles

	// SQL query to select all users
	query := "SELECT id, name,description FROM plak_roles"

	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan each row into a PlakUser struct
	for rows.Next() {
		var role views.Plak3Roles
		if err := rows.Scan(&role.Id, &role.Name, &role.Description); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *Plak3RolesRepository) Find(idparam string) (views.Plak3Roles, error) {
	var user views.Plak3Roles

	// SQL query to select the user by ID
	query := "SELECT id, name,description FROM plak_roles WHERE id = $1"

	// Execute the query
	row := r.db.QueryRow(query, idparam)
	err := row.Scan(&user.Id, &user.Name, &user.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			// Optional: return a custom error indicating that no user was found
			return views.Plak3Roles{}, err
		}
		// Return the error encountered during query execution or scanning
		return views.Plak3Roles{}, err
	}

	return user, nil
}

func (r *Plak3RolesRepository) Search(role_ids []string) ([]views.Plak3Roles, error) {
	var roles []views.Plak3Roles

	// Prepare a list of placeholders for the query (e.g., $1, $2, $3)
	placeholders := make([]string, len(role_ids))
	args := make([]interface{}, len(role_ids))
	for i, id := range role_ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}
	placeholdersStr := strings.Join(placeholders, ", ")

	// SQL query to select users with IDs in the userIDs slice
	query := fmt.Sprintf("SELECT id, name,description FROM plak_roles WHERE id IN (%s)", placeholdersStr)

	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan each row into a PlakUser struct
	for rows.Next() {
		var role views.Plak3Roles
		if err := rows.Scan(&role.Id, &role.Name, &role.Description); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
