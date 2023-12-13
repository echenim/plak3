package views

type PlakUserOnBoarding struct {
	ID        int64  `json:"Id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email"`

	Password       string `json:"password"`
	PhoneNumber    string `json:"phoneNumber"`
}
