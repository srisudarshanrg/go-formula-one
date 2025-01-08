package handlers

import (
	"log"
	"net/http"

	"github.com/srisudarshanrg/go-formula-one/server/functions"
	"github.com/srisudarshanrg/go-formula-one/server/models"
	"github.com/srisudarshanrg/go-formula-one/server/render"
)

var data = make(map[string]interface{})

// Home is the handler for the home page
func (app *HandlerRepository) Home(w http.ResponseWriter, r *http.Request) {
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
	driversChampionships, err := functions.GetDriversByAchievements(app.AppConfig.Database, "championships")
	if err != nil {
		log.Println(err)
	}

	driversWins, err := functions.GetDriversByAchievements(app.AppConfig.Database, "wins")
	if err != nil {
		log.Println(err)
	}

	driversPodiums, err := functions.GetDriversByAchievements(app.AppConfig.Database, "podiums")
	if err != nil {
		log.Println(err)
	}

	driversPoles, err := functions.GetDriversByAchievements(app.AppConfig.Database, "poles")
	if err != nil {
		log.Println(err)
	}

	currentDrivers, err := functions.GetCurrentDrivers(app.AppConfig.Database)
	if err != nil {
		log.Println(err)
	}

	data["driversChampionships"] = driversChampionships
	data["driversWins"] = driversWins
	data["driversPodiums"] = driversPodiums
	data["driversPoles"] = driversPoles
	data["currentDrivers"] = currentDrivers

	err = render.RenderTemplate(w, "drivers.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}

func (app *HandlerRepository) Teams(w http.ResponseWriter, r *http.Request) {
	allTeamsChampionships, err := functions.GetAllTeamsByNumberAchievements(app.AppConfig.Database, "championships")
	if err != nil {
		log.Println(err)
	}

	allTeamsWins, err := functions.GetAllTeamsByNumberAchievements(app.AppConfig.Database, "wins")
	if err != nil {
		log.Println(err)
	}

	allTeamsPodiums, err := functions.GetAllTeamsByNumberAchievements(app.AppConfig.Database, "podiums")
	if err != nil {
		log.Println(err)
	}

	allTeamsPoles, err := functions.GetAllTeamsByNumberAchievements(app.AppConfig.Database, "poles")
	if err != nil {
		log.Println(err)
	}

	currentTeams, err := functions.GetCurrentTeams(app.AppConfig.Database)
	if err != nil {
		log.Println(err)
	}

	data["allTeamsChampionships"] = allTeamsChampionships
	data["allTeamsWins"] = allTeamsWins
	data["allTeamsPodiums"] = allTeamsPodiums
	data["allTeamsPoles"] = allTeamsPoles
	data["currentTeams"] = currentTeams

	err = render.RenderTemplate(w, "teams.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}
