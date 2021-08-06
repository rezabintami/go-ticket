package main

import (
	"net/http"
	"ticketing/controller/cTopup"
	"ticketing/controller/cUsers"
)

func main() {
	http.HandleFunc("/topup", cTopup.PostTopUpPayment)
	http.HandleFunc("/users", cUsers.GetProfile)
	http.HandleFunc("/users", cUsers.UpdateProfile)
	http.HandleFunc("/login", cUsers.Login)
	http.ListenAndServe(":8000", nil)
}
