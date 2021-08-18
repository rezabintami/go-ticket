package routes

import (

	// _middleware "ticketing/app/middleware"
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


func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")
	// f, err := os.OpenFile("testlogfile.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Errorf("cannot open 'testlogfile', (%s)", err.Error())
	// 	flag.Usage()
	// 	os.Exit(-1)
	// }
	// log.SetOutput(f)
	// apiV1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout}))

	

	// defer f.Close()
	//! TOPUP
	apiV1.POST("/topup", cl.TopUpController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
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
	apiV1.POST("/register", cl.UserController.Register)
	apiV1.POST("/login", cl.UserController.Login)
}
