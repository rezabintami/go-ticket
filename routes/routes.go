package routes

import (
	"ticketing/controllers/cAuth"
	"ticketing/controllers/cTheater"
	// "ticketing/controllers/cMovie"
	// "ticketing/controllers/cTheater"
	// "ticketing/controllers/cTickets"
	// "ticketing/controllers/cTopup"
	// "ticketing/controllers/cUsers"

	"github.com/labstack/echo/v4"
)

func InitServer() *echo.Echo {
	router := echo.New()
	// //! TOPUP
	// router.POST("/topup", cTopup.PostTopUpPayment)

	// //! USERS
	// router.GET("/users", cUsers.GetProfile)
	// router.PUT("/users", cUsers.UpdateProfile)

	// //! TICKETS
	// router.POST("/tickets", cTickets.PostTicket)
	// router.DELETE("/tickets", cTickets.CancelTicket)

	// //! MOVIE
	// router.GET("/movies", cMovie.GetMovies)

	// //! THEATER
	router.POST("/theater", cTheater.CreateTheater)
	router.GET("/theater", cTheater.GetTheater)

	//! AUTH
	router.POST("/login", cAuth.Login)
	router.POST("/register", cAuth.Register)
	return router
}
