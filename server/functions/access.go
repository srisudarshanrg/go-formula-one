package functions

import (
	"github.com/srisudarshanrg/go-formula-one/server/config"
)

var appConfig *config.AppConfig

// AppConfigAccessFunctions provides the functions package with access to the app config
func AppConfigAccessFunctions(a *config.AppConfig) {
	appConfig = a
}
