package views

// PlakUser represents a user in the system
// @Description User model.
// @Name PlakUser
type PlakUser struct {
	ID        int64  `json:"Id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email"`

	Password       string `json:"password"`
	PhoneNumber    string `json:"phoneNumber"`
	LockoutEnabled bool   `json:"lockoutEnabled"`

	Roles []string `json:"roles,omitempty"`
}

// PlakViewUser represents a user in the system
// @Description User model.
// @Name PlakViewUser
type PlakViewUser struct {
	ID        int64        `json:"Id"`
	FirstName string       `json:"firstName,omitempty"`
	LastName  string       `json:"lastName,omitempty"`
	Email     string       `json:"email"`
	Roles     []Plak3Roles `json:"roles,omitempty"`
}

type Plak3UserObject struct {
	Status bool     `json:"status,omitempty"`
	User   PlakUser `json:"user,omitempty"`
}

type Plak3UserCollection struct {
	Status bool       `json:"status,omitempty"`
	Users  []PlakUser `json:"user,omitempty"`
}
