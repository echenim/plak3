package repositories

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/plak3com/plak3/internal/models/views"
	"golang.org/x/crypto/bcrypt"
)

type Plak3UserOnBoardingRepository struct {
	db *sql.DB
}

func NewPlak3UserOnBoardingRepository(_db *sql.DB) *Plak3UserOnBoardingRepository {
	return &Plak3UserOnBoardingRepository{db: _db}
}

// public functions
func (r *Plak3UserOnBoardingRepository) UserRegistrations(u views.PlakUserOnBoarding) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	return tx.Commit()
}

// Private functions

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

func createUserSignIn(tx *sql.Tx, user_id int64, us views.PlakUser) error {
	passwordHashed, err := hashPassword(us.Password)
	if err != nil {
		return err
	}

	// SQL statement to insert a new user
	query := `INSERT INTO plak_user_sign_in (user_id,user_name,normalized_user_name,email,normalized_email,email_confirmed,
        password_hash,security_stamp,concurrency_stamp,phone_number,phone_number_confirmed,
        two_factor_enabled,lockout_enabled) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12,$13)`

	securityStamp, concurrencyStamp := generateNewStamps()

	// Use db.Exec to execute the query without returning the id
	_, err = tx.Exec(query,
		user_id,
		us.Email,
		strings.ToUpper(us.Email),
		us.Email,
		strings.ToUpper(us.Email),
		false,
		passwordHashed,
		securityStamp,
		concurrencyStamp,
		us.PhoneNumber,
		false,
		false,
		us.LockoutEnabled,
	)

	if err != nil {
		return err
	}

	return nil
}

func findUserRoles(db *sql.DB, user_id int64) []views.Plak3Roles {
	var results []views.Plak3Roles

	// SQL query to select  by userID
	query := "SELECT role_id, role_name,role_description FROM view_plak_user_roles WHERE user_id = " + strconv.Itoa(int(user_id))

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	// Iterate over the rows and scan each row into a PlakUser struct
	for rows.Next() {
		var role views.Plak3Roles
		if err := rows.Scan(&role.Id, role.Name, role.Description); err != nil {
			return nil
		}
		results = append(results, role)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil
	}

	return results
}

func generateNewStamps() (string, string) {
	newSecurityStamp := uuid.New().String()
	newConcurrencyStamp := uuid.New().String()

	return newSecurityStamp, newConcurrencyStamp
}

// HashPassword hashes given password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares raw password with it's hashed version
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
