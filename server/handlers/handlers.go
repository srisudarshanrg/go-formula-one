package handlers

import (
	"log"
	"net/http"

	"github.com/srisudarshanrg/go-formula-one/server/functions"
	"github.com/srisudarshanrg/go-formula-one/server/models"
	"github.com/srisudarshanrg/go-formula-one/server/render"
)

// Home is the handler for the home page
func (app *HandlerRepository) Home(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]interface{})

	// current teams
	currentTeams, err := functions.GetCurrentTeams(app.AppConfig.Database)
	if err != nil {
		log.Println(err)
	}
	// current drivers
	currentDrivers, err := functions.GetCurrentDrivers(app.AppConfig.Database)
	if err != nil {
		log.Println(err)
	}
	// tracks
	currentTracks, err := functions.GetCurrentTracks(app.AppConfig.Database)
	if err != nil {
		log.Println(err)
	}

	data["currentDrivers"] = currentDrivers
	data["currentTeams"] = currentTeams
	data["currentTracks"] = currentTracks

	err = render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

// Drivers is the handler for the drivers page
func (app *HandlerRepository) Drivers(w http.ResponseWriter, r *http.Request) {
	err := render.RenderTemplate(w, "drivers.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Println(err)
	}
}
