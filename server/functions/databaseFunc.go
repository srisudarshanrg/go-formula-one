package functions

import (
	"database/sql"
	"fmt"

	"github.com/srisudarshanrg/go-formula-one/server/models"
)

// GetDriversByAchievementsNumber gets those drivers with 1 or more {achievementName} e.g championships, wins and returns them as a slice in descending order
func GetDriversByAchievements(db *sql.DB, achievementName string) ([]models.Driver, error) {
	getDriversQuery := fmt.Sprintf("select * from drivers where %s > 0 order by %s desc", achievementName, achievementName)
	rows, err := db.Query(getDriversQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []models.Driver

	for rows.Next() {
		var (
			id, age, polePositions, podiums, wins, championships, numberYearsDriven int
			name, nationality, yearsDriven, teamsDriven                             string
			createdAt, updatedAt                                                    interface{}
		)

		err = rows.Scan(&id, &name, &age, &nationality, &polePositions, &podiums, &wins, &championships, &yearsDriven, &teamsDriven, &numberYearsDriven, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, models.Driver{
			ID:                id,
			Name:              name,
			Age:               age,
			Nationality:       nationality,
			PolePositions:     polePositions,
			Podiums:           podiums,
			Wins:              wins,
			Championships:     championships,
			YearsDriven:       yearsDriven,
			TeamsDriven:       teamsDriven,
			NumberYearsDriven: numberYearsDriven,
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
		})
	}

	return drivers, nil
}

// GetCurrentDrivers gets all the rows from the current_drivers table, which contains data only about drivers in 2024
func GetCurrentDrivers(db *sql.DB) ([]models.CurrentDrivers, error) {
	getCurrentDriversQuery := `select * from current_drivers`
	rows, err := db.Query(getCurrentDriversQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []models.CurrentDrivers

	for rows.Next() {
		var (
			id, number, position, points int
			name, team, teamColor        string
			percentagePoints             float64
			championshipWinner           bool
			createdAt, updatedAt         interface{}
		)

		err = rows.Scan(&id, &name, &number, &position, &points, &team, &teamColor, &percentagePoints, &championshipWinner, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, models.CurrentDrivers{
			ID:                 id,
			Name:               name,
			Number:             number,
			Position:           position,
			Points:             points,
			Team:               team,
			TeamColor:          teamColor,
			PercentagePoints:   percentagePoints,
			ChampionshipWinner: championshipWinner,
			CreatedAt:          createdAt,
			UpdatedAt:          updatedAt,
		})
	}

	return drivers, nil
}

// GetCurrentTeams gets all the rows from the current_teams table, which contains data only about teams in 2024
func GetCurrentTeams(db *sql.DB) ([]models.CurrentTeams, error) {
	getCurrentTeamsQuery := `select * from current_teams`
	rows, err := db.Query(getCurrentTeamsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currentTeams []models.CurrentTeams

	for rows.Next() {
		var (
			id, totalPoints, constructorsPosition, highestPointsHaul int
			name, drivers                                            string
			championshipWinner                                       bool
			createdAt, updatedAt                                     interface{}
		)

		err = rows.Scan(&id, &name, &drivers, &totalPoints, &constructorsPosition, &highestPointsHaul, &championshipWinner, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		currentTeams = append(currentTeams, models.CurrentTeams{
			ID:                   id,
			Name:                 name,
			Drivers:              drivers,
			TotalPoints:          totalPoints,
			ConstructorsPosition: constructorsPosition,
			HighestPointsHaul:    highestPointsHaul,
			ChampionshipWinner:   championshipWinner,
			CreatedAt:            createdAt,
			UpdatedAt:            updatedAt,
		})
	}

	return currentTeams, nil
}

// GetAllTeamsByYearRange gets those rows from the all teams table whose date of joining formula one lies in a specified date range
func GetAllTeamsByYearRange(db *sql.DB, yearStart, yearEnd int) ([]models.AllTeams, error) {
	getAllTeamsByYearRangeQuery := `select * from all_teams where year_joined > $1 and year_joined < $2`
	rows, err := db.Query(getAllTeamsByYearRangeQuery, yearStart, yearEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allTeams []models.AllTeams

	for rows.Next() {
		var (
			id, yearJoined, poles, podiums, wins, championships        int
			name, nationality, notableCars, championshipWinningDrivers string
			createdAt, updatedAt                                       interface{}
		)

		err = rows.Scan(&id, &name, &nationality, &yearJoined, &poles, &podiums, &wins, &championships, &notableCars, &championshipWinningDrivers, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		allTeams = append(allTeams, models.AllTeams{
			ID:                         id,
			Name:                       name,
			Nationality:                nationality,
			YearJoined:                 yearJoined,
			Poles:                      poles,
			Podiums:                    podiums,
			Wins:                       wins,
			Championships:              championships,
			NotableCars:                notableCars,
			ChampionshipWinningDrivers: championshipWinningDrivers,
			CreatedAt:                  createdAt,
			UpdatedAt:                  updatedAt,
		})
	}

	return allTeams, nil
}

// GetAllTeamsByYearRange gets those rows from the all teams table whose date of joining formula one lies in a specified date range
func GetAllTeamsByNumberAchievements(db *sql.DB, achievementName string) ([]models.AllTeams, error) {
	getAllTeamsByYearRangeQuery := fmt.Sprintf("select * from all_teams where %s > 0 order by %s desc", achievementName, achievementName)
	rows, err := db.Query(getAllTeamsByYearRangeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allTeams []models.AllTeams

	for rows.Next() {
		var (
			id, yearJoined, poles, podiums, wins, championships        int
			name, nationality, notableCars, championshipWinningDrivers string
			createdAt, updatedAt                                       interface{}
		)

		err = rows.Scan(&id, &name, &nationality, &yearJoined, &poles, &podiums, &wins, &championships, &notableCars, &championshipWinningDrivers, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		allTeams = append(allTeams, models.AllTeams{
			ID:                         id,
			Name:                       name,
			Nationality:                nationality,
			YearJoined:                 yearJoined,
			Poles:                      poles,
			Podiums:                    podiums,
			Wins:                       wins,
			Championships:              championships,
			NotableCars:                notableCars,
			ChampionshipWinningDrivers: championshipWinningDrivers,
			CreatedAt:                  createdAt,
			UpdatedAt:                  updatedAt,
		})
	}

	return allTeams, nil
}
