package geocodexyz

//
// Note the init() below which will populate
// the configuration pointer through the setup()
// func.
//

import (
	"net/http"
	"os"
	"strings"
)

// hold the variables that we accept, mostly useful for tests
const (
	environmentAPIKey     = "GEOCODEXYZ_API_KEY"
	environmentSendMethod = "GEOCODEXYZ_SEND_METHOD"
)

// configOptions is a general configuration struct
type configOptions struct {
	APIKey     string
	SendMethod string // either GET or POST; default to POST
}

// config is the global configuration, configured at startup. Do not directly modify this outside
// of testing or specific cases
var config *configOptions

func setup() {
	// if already set up, just return
	if config != nil {
		return
	}
	config = &configOptions{}

	config.APIKey = osHelper(environmentAPIKey, "")
	config.SendMethod = strings.ToUpper(osHelper(environmentSendMethod, http.MethodGet))
}

func init() {
	setup()
}

// osHelper is a quick helper to look at the environment for a value and, if it isn't there,
// provide a default value
func osHelper(key, defaultValue string) string {
	found := os.Getenv(key)
	if found == "" {
		return defaultValue
	}
	return found
}
