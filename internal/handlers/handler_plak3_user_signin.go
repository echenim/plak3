package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/plak3com/plak3/internal/models/entities"
	"github.com/plak3com/plak3/internal/models/views"
	"github.com/plak3com/plak3/internal/services"
	"github.com/valyala/fasthttp"
)

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
		ctx.Error("Bad Request", fasthttp.StatusBadRequest)
		return
	}
	fmt.Printf("\n====\n login information 1 : %v\n", login_param)
	user, err := s.svc.SignIn(login_param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		ctx.Error("Invalid login", fasthttp.StatusBadRequest)
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

func (s *Plak3UserSignInHandlers) Registration(ctx *fasthttp.RequestCtx) {
	user := entities.Plak3UserSignIn{}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}
