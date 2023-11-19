package main

import (
	"github.com/plak3com/plak3/internal/handlers"
	"github.com/plak3com/plak3/internal/routes"
	"github.com/plak3com/plak3/internal/services"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()

	// Provide the database connection

	// Provide services
	container.Provide(services.NewPlak3UserService)
	container.Provide(services.NewPlak3UserSignInService)

	// Provide handler
	container.Provide(handlers.NewPlak3UserHandlers)
	container.Provide(handlers.NewPlak3UserSignInHandlers)

	// Provide routes
	container.Provide(routes.NewUSerRoutes)

	return container
}
