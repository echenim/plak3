package entities

import "time"

type StartOnboarding struct {
	Id             int64     `json:"id"`
	Email          string    `json:"email"`
	EmailConfirmed bool      `json:"email_confirmed"`
	CreatedAt      time.Time `json:"created_at"`
}
