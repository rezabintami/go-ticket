package main

import (
	"net/http"
	"ticketing/controller/cAuth"
	"ticketing/controller/cMovie"
	"ticketing/controller/cTheater"
	"ticketing/controller/cTickets"
	"ticketing/controller/cTopup"
	"ticketing/controller/cUsers"
)

func main() {
	//! TOPUP
	http.HandleFunc("/topup", cTopup.PostTopUpPayment)

	//! USERS
	http.HandleFunc("/users", cUsers.GetProfile)
	http.HandleFunc("/users", cUsers.UpdateProfile)

	//! TICKETS
	http.HandleFunc("/tickets", cTickets.PostTicket)
	http.HandleFunc("/tickets", cTickets.CancelTicket)

	//! MOVIE
	http.HandleFunc("/movie", cMovie.GetMovies)

	//! THEATER
	http.HandleFunc("/theater", cTheater.CreateTheater)
	http.HandleFunc("/theater", cTheater.GetTheater)

	//! AUTH
	http.HandleFunc("/login", cAuth.Login)
	http.HandleFunc("/register", cAuth.Register)

	http.ListenAndServe(":8000", nil)
}
