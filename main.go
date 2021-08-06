package main

import (
	"net/http"
	"ticketing/controller/cTopup"
)

func main() {
	http.HandleFunc("/topup", cTopup.PostTopUpPayment)
	http.ListenAndServe(":8000", nil)
}
