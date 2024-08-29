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
	router.GET("/", tracker.GetFingerprintHandler)

	handler := cors.New(cors.Options{
		AllowedOrigins:   config.Cors.AllowedOrigins,
		AllowCredentials: config.Cors.AllowCredentials,
	}).Handler(router)

	log.Println("Starting HTTP server on:", config.Tracker.GetAddr())
	log.Fatal(http.ListenAndServe(config.Tracker.GetAddr(), handler))
}
