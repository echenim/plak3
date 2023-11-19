package routes

import (
	"github.com/fasthttp/router"
	"github.com/plak3com/plak3/internal/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type Server struct {
	userHandler       handlers.Plak3UserHandlers
	userSigninHandler handlers.Plak3UserSignInHandlers
}

func NewUSerRoutes(_userHandler handlers.Plak3UserHandlers,
	_userSigninHandler handlers.Plak3UserSignInHandlers,
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

	// Register all Get methods
	r.GET("/api/v1/user", s.userHandler.List)

	// Register find by id methods
	r.GET("/api/v1/user/:id", s.userHandler.Find)

	// Register all Post methods
	r.POST("/api/v1/user/add", s.userHandler.Create)
	r.POST("/api/v1/user/search", s.userHandler.Search)
	r.POST("/api/v1/user/signin", s.userSigninHandler.SignIn)

	// Register all Update methods
	r.POST("/api/v1/user/edit", s.userHandler.Update)

	// Register all Delete methods
	r.DELETE("/api/v1/user/remove", s.userHandler.Remove)

	return r
}
