package handlers

import (
	"encoding/json"

	"github.com/plak3com/plak3/internal/models/entities"
	"github.com/plak3com/plak3/internal/services"
	"github.com/valyala/fasthttp"
)

type Plak3UserSignInHandlers struct {
	svc *services.Plak3UserSignInService
}

func NewPlak3UserSignInHandlers(_svc *services.Plak3UserSignInService) *Plak3UserSignInHandlers {
	return &Plak3UserSignInHandlers{svc: _svc}
}

func (h *Plak3UserSignInHandlers) SignIn(ctx *fasthttp.RequestCtx) {
	user := entities.Plak3UserSignIn{}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

func (h *Plak3UserSignInHandlers) RevokeUser(ctx *fasthttp.RequestCtx) {
	user := entities.Plak3UserSignIn{}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

func (h *Plak3UserSignInHandlers) RemoveUser(ctx *fasthttp.RequestCtx) {
	user := entities.Plak3UserSignIn{}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

func (h *Plak3UserSignInHandlers) CreateUserSignIn(ctx *fasthttp.RequestCtx) {
	user := entities.Plak3UserSignIn{}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}
