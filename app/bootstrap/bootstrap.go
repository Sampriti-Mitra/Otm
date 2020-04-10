// Package bootstrap - initialize all the components required for the application to start
// This will set the application environments and sets default logging for framework
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"otm/app/config"
	"otm/app/controllers"
	"otm/app/providers/database"

	//"otm/app/providers/database"
	"otm/app/routing/router"
	//"otm/app/validators"
)

var initializeRouter = router.Initialize
var initializeDatabase = database.Initialize

//var initializeLogger = logger.InitLogrus
var loadConfig = config.LoadConfig
var newAppController = controllers.NewAppController

//var newRuleEngineController = controllers.NewRuleEngineController
//var newClientController = controllers.NewClientController
//var newNamespaceController = controllers.NewNamespaceController
//var NewRuleChainController = controllers.NewRuleChainController
//var NewRuleGroupController = controllers.NewRuleGroupController
//var NewRuleController = controllers.NewRuleController
//var NewTemplateController = controllers.NewTemplateController
//var RegisterValidators = validators.RegisterValidators
var setGinMode = gin.SetMode

// Initialize : initializes all required application components
func Initialize(basePath string, env string) {
	baseInit(basePath, env)
	initializeRouter()
}

// BaseInit : Basic initializations required for tests
func baseInit(basePath string, env string) {
	loadConfig(basePath, env)
	setEnvironment()
	initProviders(basePath)
	initializeRequestHandlers()
}

// initProviders : Provider initialization is done here
// There initiated providers will be available across the application
func initProviders(basePath string) {
	initializeDatabase()
}

// initializeRequestHandlers : initializing request handlers
func initializeRequestHandlers() {
	newAppController()
	//newStatusController()
	//newRuleEngineController()
	//newClientController()
	//newNamespaceController()
	//NewRuleChainController()
	//NewRuleGroupController()
	//NewRuleController()
	//NewTemplateController()
	//RegisterValidators()
}

// setEnvironment : sets application gin environment based on application mode
// gin default log writer will be changed for the `release` mode
func setEnvironment() {
	appMode := config.GetConfig().Application.Mode
	setGinMode(appMode)

	if appMode == gin.ReleaseMode {
		// Disabling gin logs for release mode
		gin.DefaultWriter = ioutil.Discard
	}
}
