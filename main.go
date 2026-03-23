package main

import (
	"log"
	"net/http"

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

	log.Fatal(http.ListenAndServe(":"+app.Config.Port, nil))
}
