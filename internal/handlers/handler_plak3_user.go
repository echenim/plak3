package handlers

import (
	"encoding/json"

	"github.com/plak3com/plak3/internal/models"
	"github.com/plak3com/plak3/internal/models/searchmodels"
	"github.com/plak3com/plak3/internal/services"
	"github.com/valyala/fasthttp"
)

type Plak3UserHandlers struct {
	svc services.Plak3UserService
}

func NewPlak3UserHandlers(_svc services.Plak3UserService) *Plak3UserHandlers {
	return &Plak3UserHandlers{svc: _svc}
}

// List gets a list of users.
// @Summary List users
// @Description Retrieves a list of users.
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.PlakUser
// @Router /users [get]
func (s *Plak3UserHandlers) List(ctx *fasthttp.RequestCtx) {
	user := []models.PlakUser{}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

// Find gets a user.
// @Summary Get a user
// @Description Get details of a specific user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PlakUser
// @Router /user/find [get]
func (s *Plak3UserHandlers) Find(ctx *fasthttp.RequestCtx) {
	user := models.PlakUser{}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

// Search finds users based on criteria.
// @Summary Search users
// @Description Search for users based on criteria
// @Tags users
// @Accept  json
// @Produce  json
// @Param   criteria  body    searchmodels.UserSearchCriteria true  "Search Criteria"
// @Success 200 {array} models.PlakUser
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /user/search [post]
func (s *Plak3UserHandlers) Search(ctx *fasthttp.RequestCtx) {
	var criteria searchmodels.UserSearchCriteria

	// Assuming the criteria is sent as a JSON body
	if err := json.Unmarshal(ctx.PostBody(), &criteria); err != nil {
		ctx.Error("Bad Request", fasthttp.StatusBadRequest)
		return
	}

	users, err := s.svc.Search(criteria)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(users)
}

// Create adds a new user.
// @Summary Add a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 201 {object} models.PlakUser
// @Router /user/create [post]
func (s *Plak3UserHandlers) Create(ctx *fasthttp.RequestCtx) {
	user := models.PlakUser{}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

// Update modifies a user.
// @Summary Update a user
// @Description Update an existing user's details
// @Tags users
// @Accept  json
// @Produce  json
// @Success 201 {string} string "Updated"
// @Router /user/update [put]
func (s *Plak3UserHandlers) Update(ctx *fasthttp.RequestCtx) {
	user := models.PlakUser{}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}

// Remove deletes a user.
// @Summary Delete a user
// @Description Delete a user's account
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Deleted"
// @Router /user/remove [delete]
func (r *Plak3UserHandlers) Remove(ctx *fasthttp.RequestCtx) {
	user := models.PlakUser{}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(user)
}
