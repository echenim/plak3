package main

import (
	"log"

	"github.com/plak3com/plak3/cmd/docs"
	"github.com/plak3com/plak3/internal/routes"
	"github.com/valyala/fasthttp"
)

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	setupSwagger()

	initializeConfig()

	container := buildContainer()

	// Invoke the server start using the dependencies
	err := container.Invoke(func(server *routes.Server) {
		r := server.SetUpRoutes()
		// log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
		log.Println("Server is running at http://localhost:8080")
		if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
			panic(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}

func setupSwagger() {
	docs.SwaggerInfo.Title = "Freight Management System API Documentation"
	docs.SwaggerInfo.Description = "Freight Management System (FMS) API provides a suite of digital tools to efficiently manage and track freight operations."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
