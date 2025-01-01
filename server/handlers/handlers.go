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

	// drivers
	driversByChampionships, err := functions.GetDriversByAchievements(app.AppConfig.Database, "championships")
	if err != nil {
		log.Println(err)
	}
	driversByWins, err := functions.GetDriversByAchievements(app.AppConfig.Database, "wins")
	if err != nil {
		log.Println(err)
	}
	driversByPodiums, err := functions.GetDriversByAchievements(app.AppConfig.Database, "podiums")
	if err != nil {
		log.Println(err)
	}
	driversByPoles, err := functions.GetDriversByAchievements(app.AppConfig.Database, "pole_positions")
	if err != nil {
		log.Println(err)
	}

	// all teams
	allTeamsByChampionships, err := functions.GetAllTeamsByNumberAchievements(app.AppConfig.Database, "championships")
	if err != nil {
		log.Println(err)
	}
	allTeamsByWins, err := functions.GetAllTeamsByNumberAchievements(app.AppConfig.Database, "wins")
	if err != nil {
		log.Println(err)
	}
	allTeamsByPodiums, err := functions.GetAllTeamsByNumberAchievements(app.AppConfig.Database, "podiums")
	if err != nil {
		log.Println(err)
	}
	allTeamsByPoles, err := functions.GetAllTeamsByNumberAchievements(app.AppConfig.Database, "poles")
	if err != nil {
		log.Println(err)
	}

	data["driverChampionships"] = driversByChampionships
	data["driverWins"] = driversByWins
	data["driverPodiums"] = driversByPodiums
	data["driverPoles"] = driversByPoles

	data["currentTeams"] = currentTeams

	data["allTeamsChampionships"] = allTeamsByChampionships
	data["allTeamsWins"] = allTeamsByWins
	data["allTeamsPodiums"] = allTeamsByPodiums
	data["allTeamsPoles"] = allTeamsByPoles

	err = render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
	if err != nil {
		log.Println(err)
	}
}
