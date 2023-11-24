package main

import (
	"database/sql"
	"log"

	"github.com/plak3com/plak3/internal/handlers"
	"github.com/plak3com/plak3/internal/repositories"
	"github.com/plak3com/plak3/internal/routes"
	"github.com/plak3com/plak3/internal/services"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

func initializeConfig() {
	viper.SetConfigName("config/config") // name of config file (without extension)
	viper.SetConfigType("yaml")          // or viper.SetConfigType("json")
	viper.AddConfigPath(".")             // optionally look for config in the working directory
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	// Provide the database connection
	// container.Provide(repositories.InitDatabase(viper.GetString("database.driver"), viper.GetString("database.dataSourceName")))
	container.Provide(func() (*sql.DB, error) {
		return repositories.InitDatabase(viper.GetString("database.driver"), viper.GetString("database.dataSourceName"))
	})

	// Provide repositories
	container.Provide(repositories.NewPlak3UserRepository)
	container.Provide(repositories.NewPlak3UserSignInRepository)

	// Provide services
	container.Provide(services.NewPlak3UserService)
	container.Provide(services.NewPlak3UserSignInService)

	// Provide handler
	container.Provide(handlers.NewPlak3UserHandlers)
	container.Provide(handlers.NewPlak3UserSignInHandlers)

	// Provide routes
	container.Provide(routes.NewRoutes)

	return container
}
