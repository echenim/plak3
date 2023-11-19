package main

import (
	"log"

	"github.com/plak3com/plak3/cmd/app/docs"
	"github.com/plak3com/plak3/internal/routes"
	"github.com/valyala/fasthttp"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
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
	docs.SwaggerInfo.Description = "This is a sample server using fasthttp and swagger"
	docs.SwaggerInfo.Version = "1.0.0-0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
