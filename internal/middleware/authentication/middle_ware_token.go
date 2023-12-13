package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/plak3com/plak3/internal/models/views"
	"github.com/valyala/fasthttp"
)

func VerifyToken(tokenString string) (*claims, error) {
	claims := &claims{}
	// Parse the token with claims, directly handling the error.
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Directly return if the token is not valid.
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func tokenVarifications(tokenString string) (*claims, error) {
	claims := &claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GenerateUserTokenAfterLoginWasSuccessful(u views.Plak3SignedInUser) (string, error) {
	expirationTime := time.Now().Add(TokenExpiryDuration)
	issuedAt := time.Now()

	claims := &claims{
		UserId:          u.Id,
		Name:            u.Name,
		Email:           u.Email,
		AuthorizationTo: u.AuthorizationTo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  issuedAt.Unix(),
			Issuer:    "saas-plak3.com",
			Audience:  "FreightManagementSystem",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}
	return signedToken, nil
}

func extractBearerToken(ctx *fasthttp.RequestCtx) (string, error) {
	authToken := string(ctx.Request.Header.Peek(tokenHeader))
	if authToken == "" {
		return "", fmt.Errorf("authorization token missing")
	}
	fmt.Printf("Authorization token : %v\n", authToken)
	// if !strings.HasPrefix(authToken, bearerPrefix) {
	// 	return "", fmt.Errorf("malformed authorization token")
	// }
	// return strings.TrimPrefix(authToken, bearerPrefix), nil
	return authToken, nil
}
