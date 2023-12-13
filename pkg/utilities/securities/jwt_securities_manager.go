package securities

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/plak3com/plak3/internal/models/views"
	"github.com/valyala/fasthttp"
)

var jwtKey = []byte("aa3d118b-d505-482c-85c4-3cfa47d1ef45")

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

// VerifyToken verifies the JWT token and returns the claims.
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

func AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tokenString, err := extractBearerToken(ctx)
		if err != nil {
			sendError(ctx, err.Error(), fasthttp.StatusUnauthorized)
			return
		}

		claims, err := VerifyToken(tokenString)
		if err != nil {
			sendError(ctx, "Invalid or expired token", fasthttp.StatusUnauthorized)
			return
		}

		ctx.SetUserValue(userIDKey, claims.UserId)
		next(ctx)
	}
}

func RoleBaseAuthMiddleware(allowedRole string, next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tokenString, err := extractBearerToken(ctx)
		if err != nil {
			sendError(ctx, err.Error(), fasthttp.StatusUnauthorized)
			return
		}

		claims, err := tokenVarifications(tokenString)
		if err != nil {
			ctx.Error("Unauthorized", fasthttp.StatusUnauthorized)
			return
		}

		userRole := claims.AuthorizationTo
		if !isAllowedRole(userRole, allowedRole) {
			sendError(ctx, "Forbidden - User role not authorized", fasthttp.StatusForbidden)
			return
		}

		ctx.SetUserValue(userIDKey, claims.UserId)
		next(ctx)
	}
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

func sendError(ctx *fasthttp.RequestCtx, message string, statusCode int) {
	ctx.SetStatusCode(statusCode)
	ctx.SetContentType(contentTypeJSON)
	ctx.Error(message, statusCode)
}

func isAllowedRole(userRole []string, allowedRoles string) bool {
	for _, role := range userRole {
		if strings.TrimSpace(role) == strings.TrimSpace(allowedRoles) {
			return true
		}
	}
	return false
}
