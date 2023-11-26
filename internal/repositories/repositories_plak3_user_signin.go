package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/plak3com/plak3/internal/models/views"
)

type Plak3UserSignInRepository struct {
	db *sql.DB
}

func NewPlak3UserSignInRepository(_db *sql.DB) *Plak3UserSignInRepository {
	return &Plak3UserSignInRepository{db: _db}
}

func (r *Plak3UserSignInRepository) SignIn(u views.Plak3Login) (views.Plak3LoginUser, error) {
	userz := views.Plak3LoginUser{}
	login_user_info := views.Plak3SignedInUser{}

	// SQL query to select the user by ID
	query := "SELECT user_id, name, username, email,  passwordhash, email_confirmed, phone_number_confirmed, mfa_enabled, lockedout,roles,security_stamp FROM  view_plak_user_signed_in_roles WHERE username = $1"

	// Execute SQL query
	row := r.db.QueryRow(query, u.UserName)
	var rolesJSON string

	if err := row.Scan(&login_user_info.Id, &login_user_info.Name, &login_user_info.UserName, &login_user_info.Email, &login_user_info.PasswordHash, &login_user_info.EmailConfirmed,
		&login_user_info.PhoneNumberConfirmed, &login_user_info.MfaEnabled, &login_user_info.LockedOut, &rolesJSON, &login_user_info.SecurityStamp); err != nil {

		userz.Succeeded = false
		if err == sql.ErrNoRows {
			// Optional: return a custom error indicating that no user was found
			return userz, errors.New(" invalid user login ")
		}
		// Return the error encountered during query execution or scanning
		return userz, errors.New(" error encountered during execution ")
	}

	if err := json.Unmarshal([]byte(rolesJSON), &login_user_info.AuthorizationTo); err != nil {
		userz.Succeeded = false
		if err == sql.ErrNoRows {
			// Optional: return a custom error indicating that no user was found
			return userz, errors.New(" invalid user login ")
		}
		// Return the error encountered during query execution or scanning
		return userz, errors.New(" error encountered during execution ")
	}

	// login_user_info.AuthorizationTo = &rolesJSON
	isPasswordValid := checkPasswordHash(u.Password, login_user_info.PasswordHash)
	if isPasswordValid {
		userz.SignedInUser = login_user_info
		userz.Succeeded = true
		return userz, nil
	}

	return userz, errors.New(" invalid user signin ")
}

func (r *Plak3UserSignInRepository) Revoke(ctx context.Context) error {
	return nil
}

func (r *Plak3UserSignInRepository) Remove(ctx context.Context) error {
	return nil
}

func (r *Plak3UserSignInRepository) createSignIn(ctx context.Context) error {
	return nil
}
