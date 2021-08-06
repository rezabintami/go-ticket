package main

import (
	"net/http"
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

	//! Ticket
	http.HandleFunc("/tickets", cTickets.PostTicket)
	http.HandleFunc("/tickets", cTickets.CancelTicket)

	//! AUTH
	http.HandleFunc("/login", cUsers.Login)
	http.HandleFunc("/register", cUsers.Register)

	http.ListenAndServe(":8000", nil)
}
