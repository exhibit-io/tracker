package main

import (
	"log"
	"net/http"

	"github.com/exhibit-io/tracker"
	"github.com/exhibit-io/tracker/config"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	config := config.LoadConfig()

	tracker.Init(config)

	// Use a Router to handle requests
	router := httprouter.New()
	router.GET("/", tracker.TrackerHandler)

	// Setup middlewares.  For this we're basically adding:
	//	- Support for CORS to make JSONP work.
	handler := cors.Default().Handler(router)

	log.Println("Starting HTTP server on:", config.Tracker.GetAddr())
	log.Fatal(http.ListenAndServe(config.Tracker.GetAddr(), handler))
}
