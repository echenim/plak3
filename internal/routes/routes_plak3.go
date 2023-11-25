package routes

import (
	"github.com/fasthttp/router"
	"github.com/plak3com/plak3/internal/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type Server struct {
	userHandler       *handlers.Plak3UserHandlers
	userSigninHandler *handlers.Plak3UserSignInHandlers
}

func NewRoutes(_userHandler *handlers.Plak3UserHandlers,
	_userSigninHandler *handlers.Plak3UserSignInHandlers,
) *Server {
	return &Server{
		userHandler:       _userHandler,
		userSigninHandler: _userSigninHandler,
	}
}

func (s *Server) SetUpRoutes() *router.Router {
	r := router.New()

	// Register Swagger API routes
	r.GET("/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))

	// user login managment routes
	r.POST("/api/v1/login", s.userSigninHandler.LoginIn)

	// user management routes
	r.GET("/api/v1/users", s.userHandler.List)
	r.POST("/api/v1/users", s.userHandler.Create)
	r.GET("/api/v1/users/{id}", s.userHandler.Find)
	r.PUT("/api/v1/users/{id}", s.userHandler.Update)
	r.DELETE("/api/v1/users/{id}", s.userHandler.Remove)
	r.POST("/api/v1/users/search", s.userHandler.Search)

	// Register all Delete methods

	return r
}
