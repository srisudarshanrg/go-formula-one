package validations

import "github.com/srisudarshanrg/go-formula-one/server/config"

var appConfig *config.AppConfig

// AppConfigAccessValidations provides the validations package with access to the app config
func AppConfigAccessValidations(a *config.AppConfig) {
	appConfig = a
}
