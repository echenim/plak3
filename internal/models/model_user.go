package models

// User represents a user in the system.
// swagger:model User
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
