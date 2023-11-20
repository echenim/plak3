package entities

import "time"

type Plak3UserSignIn struct {
	ID                       string     `json:"id"`
	UserId                   int64      `json:"user_id"`
	UserName                 string     `json:"userName"`
	NormalizedUserName       string     `json:"normalizedUserName"`
	Email                    string     `json:"email"`
	NormalizedEmail          string     `json:"normalizedEmail"`
	EmailConfirmed           bool       `json:"emailConfirmed"`
	PasswordHash             string     `json:"passwordHash"`
	SecurityStamp            string     `json:"securityStamp"`
	ConcurrencyStamp         string     `json:"concurrencyStamp"`
	PhoneNumber              string     `json:"phoneNumber"`
	PhoneNumberConfirmed     bool       `json:"phoneNumberConfirmed"`
	TwoFactorEnabled         bool       `json:"twoFactorEnabled"`
	LockoutEnd               time.Time  `json:"lockoutEnd"`
	LockoutEnabled           bool       `json:"lockoutEnabled"`
	AccessFailedCount        int        `json:"accessFailedCount"`
	LastSuccessfulSignInDate *time.Time `json:"lastSuccessFulSignInDate"`
	LastFailedSignInDate     *time.Time `json:"lastFailedSignInDate"`
}
