package main

import (
	"log"
	"net/http"
	"time"

	"app/internal/application"
	"app/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := application.New(cfg)

	http.HandleFunc("/info", app.InfoHandler)
	http.HandleFunc("/info/weather", app.WeatherHandler)

	server := &http.Server{
		Addr:        ":" + app.Config.Port,
		ReadTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
