package handlers

import (
	"github.com/srisudarshanrg/go-formula-one/server/config"
)

var Repository *HandlerRepository

// HandlerRepository is the repository for the handlers package which contains the app config, database and session
type HandlerRepository struct {
	AppConfig *config.AppConfig
}

// RepositoryAccesshandlers provides the handlers package with access to the repository
func RepositoryAccesshandlers(repoSet *HandlerRepository) {
	Repository = repoSet
}
