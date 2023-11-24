package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/plak3com/plak3/internal/models/searchmodels"
	"github.com/plak3com/plak3/internal/models/views"
	"github.com/plak3com/plak3/internal/services"
	"github.com/valyala/fasthttp"
)

type Plak3UserHandlers struct {
	svc *services.Plak3UserService
}

func NewPlak3UserHandlers(_svc *services.Plak3UserService) *Plak3UserHandlers {
	return &Plak3UserHandlers{svc: _svc}
}


// List gets a list of users.
// @Summary List users
// @Description Retrieves a list of users.
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} views.PlakUser
// @Router /users [get]
func (s *Plak3UserHandlers) List(ctx *fasthttp.RequestCtx) {
	user, err := s.svc.Get()
	if err != nil {
		// Set an appropriate error status code
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		// Optionally, you can send a JSON response with the error details
		json.NewEncoder(ctx).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Successful response
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Failed to encode user data"})
	}
}

// Find gets a user.
// @Summary Get a user
// @Description Get details of a specific user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} views.PlakUser
// @Router /user/find [get]
func (s *Plak3UserHandlers) Find(ctx *fasthttp.RequestCtx) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		// Set an appropriate error status code
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		ctx.Error("Invalid user ID", fasthttp.StatusBadRequest)
		return
	}
	user, err := s.svc.Find(int64(id))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetContentType("application/json")
		ctx.Error("User not found", fasthttp.StatusNotFound)
		return
	}
	// Successful response
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Failed to encode user data"})
	}
}

// Search finds users based on criteria.
// @Summary Search users
// @Description Search for users based on criteria
// @Tags users
// @Accept  json
// @Produce  json
// @Param   criteria  body    searchmodels.UserSearchCriteria true  "Search Criteria"
// @Success 200 {array} views.PlakUser
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /user/search [post]
func (s *Plak3UserHandlers) Search(ctx *fasthttp.RequestCtx) {
	var criteria searchmodels.UserSearchCriteria

	// Assuming the criteria is sent as a JSON body
	if err := json.Unmarshal(ctx.PostBody(), &criteria); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		ctx.Error("Bad Request", fasthttp.StatusBadRequest)
		return
	}

	users, exist, err := s.svc.Search(criteria)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		ctx.Error("Server Error", fasthttp.StatusInternalServerError)
		return
	}
	if !exist {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetContentType("application/json")
		ctx.Error("No result found", fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(users); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Failed to encode user data"})
	}
}

// Create adds a new user.
// @Summary Add a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 201 {object} views.PlakUser
// @Router /user/create [post]
func (s *Plak3UserHandlers) Create(ctx *fasthttp.RequestCtx) {
	var newUser views.PlakUser
	if err := json.Unmarshal(ctx.PostBody(), &newUser); err != nil {
		// Handle JSON unmarshal error
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Invalid user data"})
		return
	}

	// Save the user using the service
	user, err := s.svc.Save(newUser)
	if err != nil {
		// Handle error from the Save method
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Failed to save user"})
		return
	}

	// Respond with the created user
	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Failed to encode user data"})
	}
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
	// id, err := strconv.Atoi(ctx.UserValue("id").(string))
	// if err != nil {
	// 	ctx.SetStatusCode(fasthttp.StatusBadRequest)
	// 	ctx.SetContentType("application/json")
	// 	ctx.Error("Invalid user ID", fasthttp.StatusBadRequest)
	// 	return
	// }
	var updateUser views.PlakUser
	if err := json.Unmarshal(ctx.PostBody(), &updateUser); err != nil {
		// Handle JSON unmarshal error
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Invalid user data"})
		return
	}

	_, err := s.svc.Find(updateUser.ID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetContentType("application/json")
		ctx.Error("User not found", fasthttp.StatusNotFound)
		return
	}
	// Save the user using the service
	user, err := s.svc.Edit(updateUser)
	if err != nil {
		// Handle error from the Save method
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Failed to update user record"})
		return
	}

	// Respond with the created user
	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Failed to encode user data"})
	}
}

// Remove deletes a user.
// @Summary Delete a user
// @Description Delete a user's account
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Deleted"
// @Router /user/remove [delete]
func (s *Plak3UserHandlers) Remove(ctx *fasthttp.RequestCtx) {
	id, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetContentType("application/json")
		ctx.Error("Invalid user ID", fasthttp.StatusBadRequest)
		return
	}

	_, err = s.svc.Find(int64(id))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetContentType("application/json")
		ctx.Error("User not found", fasthttp.StatusNotFound)
		return
	}

	// Save the user using the service
	err = s.svc.Remove(int64(id))
	if err != nil {
		// Handle error from the Save method
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		json.NewEncoder(ctx).Encode(map[string]string{"error": "Failed to update user record"})
		return
	}

	// Respond with the delete user
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
