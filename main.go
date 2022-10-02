package main

import (
	"party_invite/controllers"
)

func main() {
	// Startup controller
	c := &controllers.Controller{}

	// Start webserver & controllers
	c.StartWebServer()
}
