package flags

import (
	"fmt"
	"globa_trotter_game/constants"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	// env            = flag.String(constants.EnvKey, constants.EnvDefaultValue, constants.EnvUsage)
	// port           = flag.Int(constants.PortKey, constants.PortDefaultValue, constants.PortUsage)
	// metricPort     = flag.Int(constants.MetricPortKey, constants.MetricPortDefaultValue, constants.PortUsage)
	// pprofPort      = flag.Int(constants.PProfPortKey, constants.PProfPortDefaultValue, constants.PortUsage)
	baseConfigPath = flag.String(constants.BaseConfigPathKey, constants.BaseConfigPathDefaultValue,
		constants.BaseConfigPathUsage)
	// appVersion = flag.String(constants.AppVersionKey, constants.AppVersionDefaultValue, constants.AppVersionUsage)
	// dc         = flag.String(constants.DCNameKey, constants.DCDefaultValue, constants.DCNameUsage)
)

func init() {
	fmt.Println("flags inits")
	flag.Parse()
}

// // Env is the application.yml runtime environment
// func Env() string {
// 	return *env
// }

// // Port is the application.yml port number where the process will be started
// func Port() int {
// 	return *port
// }

// BaseConfigPath is the path that holds the configuration files
func BaseConfigPath() string {
	return *baseConfigPath
}

// func MetricPort() int {
// 	return *metricPort
// }

// func PProfPort() int {
// 	return *pprofPort
// }

// func AppVersion() string {
// 	return *appVersion
// }
// func DCName() string {
// 	return *dc
// }

func PTRMS_TestDb_UserName() string {
	e := os.Getenv("PTRMS_TESTDB_USERNAME")
	if e == "" {
		return ""
	}
	return e
}

func PTRMS_TestDb_Password() string {
	e := os.Getenv("PTRMS_TESTDB_PASSWORD")
	if e == "" {
		return ""
	}
	return e
}
