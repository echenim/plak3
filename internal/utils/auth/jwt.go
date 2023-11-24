package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/plak3com/plak3/internal/models/viewmodels"
)

type JWTClaims struct {
	UserId             int64  `json:"user_id"`
	Roles              string `json:roles"`
	Organization       string `json:organization"`
	jwt.StandardClaims `json:standard_claims"`
}

const JwtKey = "your_secret_key" // Replace with a secure key

func GenerateToken(user viewmodels.SignedIn) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaims{
		UserId:       user.UserHasSignedIn.UserId,
		Roles:        user.UserHasSignedIn.UserRoles,
		Organization: user.UserHasSignedIn.Organization,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(JwtKey))
}

func ValidateToken(tokenString string) (*jwt.Token, *JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})

	return token, claims, err
}
