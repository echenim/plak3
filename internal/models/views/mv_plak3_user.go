package views

type PlakUser struct {
	ID        int64    `json:"Id"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles,omitempty"`
}

type Plak3UserObject struct {
	Status bool     `json:"status,omitempty"`
	User   PlakUser `json:"user,omitempty"`
}

type Plak3UserCollection struct {
	Status bool       `json:"status,omitempty"`
	Users  []PlakUser `json:"user,omitempty"`
}
