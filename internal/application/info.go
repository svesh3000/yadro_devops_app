package application

import (
	"net/http"
)

func (app *App) InfoHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"version": app.Config.Version,
		"service": app.Config.Service,
		"author":  app.Config.Author,
	}

	writeJSONResponse(w, response)
}
