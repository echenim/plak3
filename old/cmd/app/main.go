package main

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/plak3com/plak3/cmd/app/docs"
	h "github.com/plak3com/plak3/internal/handlers"
	"github.com/valyala/fasthttp"
)

func main() {
	setupSwagger()

	r := h.SetupRouter()

	startServer(r)
}

func setupSwagger() {
	docs.SwaggerInfo.Title = "Fasthttp Swagger"
	docs.SwaggerInfo.Description = "This is a sample server using fasthttp and swagger"
	docs.SwaggerInfo.Version = "1.0.0-0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func startServer(r *router.Router) {
	address := ":8080"
	log.Printf("Server is running at http://localhost%s", address)
	if err := fasthttp.ListenAndServe(address, r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
