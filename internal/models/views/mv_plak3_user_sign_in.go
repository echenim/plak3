package views

// Plak3UserSignIn represents signed in user
// @Description SingnedIn model.
// @Name Plak3UserSignIn
type Plak3UserSignIn struct {
	ID               string `json:"id"`
	UserId           int64  `json:"user_id"`
	UserName         string `json:"userName"`
	Email            string `json:"email"`
	PasswordHash     string `json:"passwordHash"`
	PhoneNumber      string `json:"phoneNumber"`
	LockoutEnabled   bool   `json:"lockoutEnabled"`
	TwoFactorEnabled bool   `json:"twoFactorEnabled"`
}

// Plak3Login represents user login
// @Description login model.
// @Name Plak3Login
type Plak3Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// Plak3SignedInUser represents user login
// @Description login response model.
// @Name Plak3SignedInUser
type Plak3SignedInUser struct {
	Id                   int64  `json:"id"`
	Name                 string `json:"name"`
	UserName             string `json:"user_name"`
	Email                string `json:"email"`
	PasswordHash         string `json:"password_hash"`
	EmailConfirmed       bool   `json:"email_confirmed"`
	PhoneNumberConfirmed bool   `json:"phone_number_confirmed"`
	MfaEnabled           bool   `json:"mfa_enabled"`
	LockedOut            bool   `json:"lockedout"`

	AuthorizationTo []Plak3Roles `json:"authorization_to,omitempty"`
}

// Plak3LoginUser represents user login
// @Description login  model.
// @Name Plak3LoginUser
type Plak3LoginUser struct {
	Succeeded    bool              `json:"succeeded"`
	SignedInUser Plak3SignedInUser `json:"signed_in_user"`
}
