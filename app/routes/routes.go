package routes

import (
	"ticketing/controllers/topup"
	"ticketing/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware   middleware.JWTConfig
	UserController  users.UserController
	TopUpController topup.TopUpController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// users := e.Group("users")

	// //! TOPUP
	e.POST("/topup", cl.TopUpController.PaymentTopUp, middleware.JWTWithConfig(cl.JWTMiddleware))

	// //! USERS
	e.GET("/users/:id", cl.UserController.GetProfile, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.PUT("/users/:id", cl.UserController.UpdateProfile, middleware.JWTWithConfig(cl.JWTMiddleware))

	// //! TICKETS
	// router.POST("/tickets", cTickets.PostTicket)
	// router.DELETE("/tickets", cTickets.CancelTicket)

	// //! MOVIE
	// e.GET("/movies", cMovie.GetMovies)

	// //! THEATER
	// e.POST("/theater", cTheater.CreateTheater)
	// e.GET("/theater", cTheater.GetTheater)

	//! AUTH
	e.POST("/register", cl.UserController.Register)
	e.POST("/login", cl.UserController.Login)

	// category := e.Group("category")
	// category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	// news := e.Group("news")
	// news.POST("/store", cl.NewsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
}
