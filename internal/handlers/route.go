package handlers

import (
	"github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func SetupRouter() *router.Router {
	r := router.New()
	r.GET("/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	r.GET("/read", readHandler)
	r.POST("/create", createUserHandler)
	return r
}
