package handlers

import (
	"encoding/json"
	"fmt"
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

// List godoc
// @Summary List all users
// @Description get users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} views.PlakViewUser
// @Failure 500 {array} views.PlakViewUser
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

// Find godoc
// @Summary Find a user by ID
// @Description Retrieves a user based on their ID.
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} views.PlakViewUser "Successfully retrieved the user"
// @Failure 400 {object} views.PlakViewUser "Invalid user ID"
// @Failure 404 {object} views.PlakViewUser "User not found"
// @Failure 500 {object} views.PlakViewUser "Internal Server Error"
// @Router /users/{id} [get]
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

// Search finds users based on search criteria
// @Summary Search for users
// @Description Searches for users based on various criteria like ID, first name, last name, and email.
// @Tags users
// @Accept json
// @Produce json
// @Param searchmodels.UserSearchCriteria body searchmodels.UserSearchCriteria true "Search Criteria"
// @Success 200 {array} views.PlakViewUser "List of users matching criteria"
// @Failure 400 {array} views.PlakViewUser "Bad Request"
// @Failure 404 {array} views.PlakViewUser "Not Found"
// @Failure 500 {array} views.PlakViewUser "Internal Server Error"
// @Router /users/search [post]
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

// Create a new user
// @Summary Create a new user
// @Description Adds a new user to the system
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body views.PlakUser true "User to create"
// @Success 201 {object} views.PlakUser "User created successfully"
// @Failure 400 {object} views.PlakUser "Invalid user data"
// @Failure 500 {object} views.PlakUser "Internal Server Error"
// @Router /users [post]
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
		fmt.Printf("\nError saving user: %v\n", err)
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

// Update updates a user's information
// @Summary Update user information
// @Description Update user information based on provided data
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body views.PlakUser true "User Data"
// @Success 200 {object} views.PlakUser "Successfully updated user"
// @Failure 400 {object} views.PlakUser "Invalid user data"
// @Failure 404 {object} views.PlakUser "User not found"
// @Failure 500 {object} views.PlakUser "Internal server error"
// @Router /users/{id} [put]
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
	user, err := s.svc.EditUser(updateUser)
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

// Remove removes a user by ID
// @Summary Remove a user
// @Description Remove a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string "message: User removed successfully"
// @Failure 400 {object} map[string]string "error: Invalid user ID"
// @Failure 404 {object} map[string]string "error: User not found"
// @Failure 500 {object} map[string]string "error: Failed to remove user"
// @Router /users/{id} [delete]
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
