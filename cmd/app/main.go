package main

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/plak3com/plak3/docs"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func main() {
	setupSwagger()

	r := setupRouter()

	startServer(r)
}

func setupSwagger() {
	docs.SwaggerInfo.Title = "Fasthttp Swagger"
	docs.SwaggerInfo.Description = "This is a sample server using fasthttp and swagger"
	docs.SwaggerInfo.Version = "1.0.0-0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func setupRouter() *router.Router {
	r := router.New()
	r.GET("/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	r.GET("/read", readHandler)
	r.POST("/create", createUserHandler)
	return r
}

func startServer(r *router.Router) {
	address := ":8080"
	log.Printf("Server is running at http://localhost%s", address)
	if err := fasthttp.ListenAndServe(address, r.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
