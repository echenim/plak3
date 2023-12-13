package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpiryDuration = 1 * time.Hour
const (
	authHeader      = "Authorization"
	bearerPrefix    = "Bearer "
	contentTypeJSON = "application/json"
	userIDKey       = "User_id"
	tokenHeader     = "Token"
)

type claims struct {
	UserId             int64    `json:"user_id"`
	Name               string   `json:"name"`
	Email              string   `json:"email"`
	AuthorizationTo    []string `json:"roles"`
	jwt.StandardClaims `json:"standard_claims"`
}
