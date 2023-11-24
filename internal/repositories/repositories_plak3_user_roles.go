package repositories

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"github.com/plak3com/plak3/internal/models/views"
)

type Plak3UserRolesRepository struct {
	db *sql.DB
}

func NewPlak3UserRolesRepository() *Plak3UserRolesRepository {
	return &Plak3UserRolesRepository{}
}

func (r *Plak3UserRolesRepository) Find(user_id int64) ([]views.Plak3UserRoles, error) {
	var results []views.Plak3UserRoles

	// SQL query to select  by userID
	query := "SELECT role_id, role_name,user_id FROM view_plak_user_roles WHERE user_id = " + strconv.Itoa(int(user_id))

	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan each row into a PlakUser struct
	for rows.Next() {
		var user_role views.Plak3UserRoles
		if err := rows.Scan(&user_role.RoleId, &user_role.RoleName, &user_role.UserId); err != nil {
			return nil, err
		}
		results = append(results, user_role)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}



func (r *Plak3UserRolesRepository) Save(u views.Plak3UserAndRoles) error {
	err := r.remove(u.UserId)
	if err != nil {
		return err
	}
	return r.save(u.UserId, u.Roles)
}

func (r *Plak3UserRolesRepository) remove(u int64) error {
	// SQL statement to delete the user
	query := "DELETE FROM plak_user_roles WHERE user_id = $1"

	// Execute the query
	_, err := r.db.Exec(query, u)
	return err
}

func (r *Plak3UserRolesRepository) save(user_id int64, u []string) error {
	// SQL statement to insert a new user
	query := `INSERT INTO  plak_user_roles (user_id,role_id) VALUES ($1, $2)`
	var errStrs []string
	for _, i := range u {
		// Use db.Exec to execute the query without returning the id
		_, err := r.db.Exec(query, user_id, i)
		if err != nil {
			errStrs = append(errStrs, err.Error())
		}
	}

	return errors.New(strings.Join(errStrs, ", "))
}
