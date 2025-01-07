package models

// Driver is the model for a driver object in the database
type Driver struct {
	ID                int
	Name              string
	Age               int
	Nationality       string
	PolePositions     int
	Podiums           int
	Wins              int
	Championships     int
	YearsDriven       string
	TeamsDriven       string
	NumberYearsDriven int
	CreatedAt         interface{}
	UpdatedAt         interface{}
}

// CurrentDrivers is the model for a current drivers object in the database
type CurrentDrivers struct {
	ID                 int
	Name               string
	Number             int
	Position           int
	Points             int
	Team               string
	TeamColor          string
	PercentagePoints   float64
	ChampionshipWinner bool
	CreatedAt          interface{}
	UpdatedAt          interface{}
}

// CurrentTeams is the model for a current teams object in the database
type CurrentTeams struct {
	ID                   int
	Name                 string
	Drivers              string
	TotalPoints          int
	ConstructorsPosition int
	HighestPointsHaul    int
	ChampionshipWinner   bool
	LogoLink             string
	CreatedAt            interface{}
	UpdatedAt            interface{}
}

// AllTeams is the model for an all teams object in the database
type AllTeams struct {
	ID                         int
	Name                       string
	Nationality                string
	YearJoined                 int
	Poles                      int
	Podiums                    int
	Wins                       int
	Championships              int
	NotableCars                string
	ChampionshipWinningDrivers string
	CreatedAt                  interface{}
	UpdatedAt                  interface{}
}

// CurrentTracks is the model for a current track object in the database
type CurrentTracks struct {
	ID              int
	Name            string
	Length          int
	NumberCorners   int
	NumberStraights int
	NumberDRSZones  int
	Year            int
	Country         string
	Image           string
	CreatedAt       interface{}
	UpdatedAt       interface{}
}
