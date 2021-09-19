package routes

import (
	"html/template"
	"io"
	_middleware "ticketing/app/middleware"
	"ticketing/controllers/movies"
	"ticketing/controllers/theater"
	"ticketing/controllers/tickets"
	"ticketing/controllers/topup"
	"ticketing/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	UserController    users.UserController
	TopUpController   topup.TopUpController
	TheaterController theater.TheaterController
	MoviesController  movies.MovieController
	TicketsController tickets.TicketsController
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/view/*.html")),
	}
	e.Renderer = t
	apiV1 := e.Group("/api/v1")

	e.Use(_middleware.MiddlewareLogging)

	//! TOPUP
	apiV1.GET("/topup/payment", cl.TopUpController.CreateTransaction, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.POST("/topup/payment/callback", cl.TopUpController.TransactionCallbackHandler, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.GET("/topup", cl.TopUpController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! USERS
	apiV1.GET("/users", cl.UserController.GetProfile, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.PUT("/users", cl.UserController.UpdateProfile, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! TICKETS
	apiV1.GET("/tickets", cl.TicketsController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.POST("/tickets", cl.TicketsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.DELETE("/tickets/:id", cl.TicketsController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! MOVIE
	apiV1.GET("/movies", cl.MoviesController.Fetch, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.GET("/movies/:id", cl.MoviesController.GetDetailMovies, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! THEATER
	apiV1.POST("/theater", cl.TheaterController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.PUT("/theater/:id", cl.TheaterController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.DELETE("/theater/:id", cl.TheaterController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	apiV1.GET("/theater", cl.TheaterController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! AUTH
	auth := apiV1.Group("/auth")
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/login", cl.UserController.Login)

	//! OAUTH2
	auth.GET("/login/oauth", cl.UserController.OauthLogin)

	//! GOOGLE
	auth.GET("/google", cl.UserController.LoginGoogle)
	auth.GET("/google/callback", cl.UserController.HandleGoogle)

	//! FACEBOOK
	auth.GET("/facebook", cl.UserController.LoginFacebook)
	auth.GET("/facebook/callback", cl.UserController.HandleFacebook)

}
