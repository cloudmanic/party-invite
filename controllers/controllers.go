package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

//
// StartWebServer - Start the webserver
//
func (t *Controller) StartWebServer() {
	// Set GIN Settings
	gin.SetMode("release")
	gin.DisableConsoleColor()

	// Set Router
	router := gin.New()

	// Logger - Global middleware
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Set a lower memory limit for multipart forms (default is 32 MiB). Good for managing DOS attacks
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Register Routes
	t.DoRoutes(router)

	// Setup http server
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	// Start server and log if fails
	log.Fatal(srv.ListenAndServe())
}
