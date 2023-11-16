package handlers

import (
	"encoding/json"

	"github.com/plak3com/plak3/internal/models"
	"github.com/valyala/fasthttp"
)

// CreateUser creates a new user.
// @Summary Create a new user
// @Description Create a new user with the given details
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User data"
// @Success 200 {object} User
// @Router /users [post]
func createUserHandler(ctx *fasthttp.RequestCtx) {
	// Here, you would parse the request body to create a new User
	// and then save this user to your database or in-memory store.
	// For the sake of this example, we'll just simulate this with a dummy response.

	var user models.User
	// Normally, you would parse and validate the request body here
	// e.g., json.Unmarshal(ctx.PostBody(), &user)

	user = models.User{ID: 1, Name: "John Doe", Email: "johndoe@example.com"} // Dummy data

	// Respond with the newly created user
	responseData, _ := json.Marshal(user)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(responseData)
}

func readHandler(ctx *fasthttp.RequestCtx) {
	// Implement read logic
}
