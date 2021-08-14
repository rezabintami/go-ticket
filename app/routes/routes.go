package routes

import (
	"ticketing/controllers/movies"
	"ticketing/controllers/theater"
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
}

//! GET MOVIES BY ID
//! GET MOVIES BY LANGUAGE
//! CREATE TICKET
//! CANCEL TICKET

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// users := e.Group("users")

	// //! TOPUP
	e.POST("/topup", cl.TopUpController.PaymentTopUp, middleware.JWTWithConfig(cl.JWTMiddleware))

	// //! USERS
	e.GET("/users", cl.UserController.GetProfile, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.PUT("/users", cl.UserController.UpdateProfile, middleware.JWTWithConfig(cl.JWTMiddleware))

	// //! TICKETS
	// router.POST("/tickets", cTickets.PostTicket)
	// router.DELETE("/tickets", cTickets.CancelTicket)

	// //! MOVIE
	e.GET("/movies", cl.MoviesController.Fetch, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("/movies/:id", cl.MoviesController.GetDetailMovies, middleware.JWTWithConfig(cl.JWTMiddleware))

	// //! THEATER
	e.POST("/theater", cl.TheaterController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.PUT("/theater/:id", cl.TheaterController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.DELETE("/theater/:id", cl.TheaterController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("/theater", cl.TheaterController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	//! AUTH
	e.POST("/register", cl.UserController.Register)
	e.POST("/login", cl.UserController.Login)

	// category := e.Group("category")
	// category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	// news := e.Group("news")
	// news.POST("/store", cl.NewsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
}
