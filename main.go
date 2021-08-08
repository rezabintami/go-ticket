package main

import (
	"ticketing/configs"
	"ticketing/routes"
)

func main() {
	configs.InitDB()
	e := routes.InitServer()
	e.Start(":8000")
}
