package handlers

import (
	"github.com/fasthttp/router"
	"github.com/plak3com/plak3/internal/handlers"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/plak3com/plak3/docs"
)

func setupRouter() *router.Router {
	r := router.New()
	r.GET("/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	r.GET("/read", readHandler)
	r.POST("/create",createUserHandler)
	return r
}
