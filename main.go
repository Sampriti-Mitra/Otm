// Application execution starts here
// This will have all blank imports used in the application as part of convention
// This will initialize the application based on the parameter specified
package main

import (
	"flag"
	"otm/app/bootstrap"
	"otm/app/constants"
)

var (
	basePath *string
	env      *string
)

func init() {
	basePath = flag.String(constants.BasePath, constants.DefaultBasePath, "Path to otm base path")
	env = flag.String(constants.Env, constants.Development, "Application env : prod/dev")
}

// Execution starts here
// while execution we can give the command line parameters to configure the application
// There are two flags which can be set while initiating
// - base_path: default it'll take the current directory as application
//			    configuration will be loaded from `conf` folder available in the directory specified by this
// - env: default it'll be `dev`. application will be initialized according to this
// - command: when specified application will initialize in command mode and execute the given command
//			  router wont be initialized for commands
func main() {
	flag.Parse()

	bootstrap.Initialize(*basePath, *env)
}
