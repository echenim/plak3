package entities

import "time"

type Plak3EmailBank struct {
	Id        string     `json:"id"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
