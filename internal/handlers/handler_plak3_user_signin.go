package handlers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/plak3com/plak3/internal/models/entities"
	"github.com/plak3com/plak3/internal/models/views"
	"github.com/plak3com/plak3/internal/services"
	"github.com/plak3com/plak3/internal/utils/securities"
	"github.com/valyala/fasthttp"
)

type JWTClaims struct {
	Name               string `json:"name"`
	UserId             int64  `json:"user_id"`
	Email              string `json:"email"`
	AuthorizationTo    string `json:"authorization_to"`
	jwt.StandardClaims `json:"standard_claims"`
}

type Plak3UserSignInHandlers struct {
	svc *services.Plak3UserSignInService
}

func NewPlak3UserSignInHandlers(_svc *services.Plak3UserSignInService) *Plak3UserSignInHandlers {
	return &Plak3UserSignInHandlers{svc: _svc}
}

// LoginIn logs in a user.
// @Summary User login
// @Description Logs in a user and returns user information.
// @Tags authenticate & authorization
// @Accept json
// @Produce json
// @Param views.Plak3Login body views.Plak3Login true "Login information"
// @Success 200 {object} views.Plak3SignedInUser "Successfully logged in"
// @Failure 400 {object} views.Plak3SignedInUser "Bad request"
// @Failure 401 {object} views.Plak3SignedInUser "Invalid login credentials"
// @Router /login [post]
func (s *Plak3UserSignInHandlers) LoginIn(ctx *fasthttp.RequestCtx) {
	var login_param views.Plak3Login

	if err := json.Unmarshal(ctx.PostBody(), &login_param); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		ctx.Error("Invalid request", fasthttp.StatusBadRequest)
		return
	}

	user, err := s.svc.SignIn(login_param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		ctx.Error("Invalid login credentials", fasthttp.StatusUnauthorized)
		return
	}

	token, err := securities.GenerateUserTokenAfterLoginWasSuccessful(user.SignedInUser)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		ctx.Error("Error generating token", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(map[string]string{"token": token})
}

func (s *Plak3UserSignInHandlers) Registration(ctx *fasthttp.RequestCtx) {
	user := entities.Plak3UserSignIn{}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

func generateUserTokenAfterLoginWasSuccessful(u views.Plak3SignedInUser) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaims{
		UserId:          u.Id,
		Name:            u.Name,
		AuthorizationTo: strings.Join(u.AuthorizationTo, " | "),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte(u.SecurityStamp))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}
	return signedToken, nil
}

func ValidateUserToken(tokenString string) (*jwt.Token, *JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	return token, claims, err
}

func generateECDSAKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}
