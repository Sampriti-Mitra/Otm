package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"otm/app/constants"
	"reflect"
)

const (
	// FilePath - relative path to the config directory
	FilePath = "%s/conf/%s"

	// DefaultFilename - Filename format of default config file
	DefaultFilename = "env.default.toml"

	// EnvFilename - Filename format of env specific config file
	EnvFilename = "env.%s.toml"
)

var (
	// config : this will hold all the application configuration
	config AppConfig
)

// appConfig global configuration struct definition
type AppConfig struct {
	Application application
	Database    DatabaseConfig `toml:"database"`
	AuthUser    authUser       `toml:"authuser"`
	//Prometheus  prometheus     `toml:"prometheus"`
	OneTouchMusic oneTouchMusic `toml:"one_touch_music"`
}

// LoadConfig will load the configuration available in the cnf directory available in basePath
// conf file will takes based on the env provided
func LoadConfig(basePath string, env string) {
	// reading conf based on default environment
	//loadConfigFromFile(basePath, DefaultFilename, "")

	// reading env file and override conf values; if env file exists
	loadConfigFromFile(basePath, EnvFilename, env)

	//Validate
	ValidateConfig("", GetConfig())
}

func ValidateConfig(field string, a interface{}) {
	v := reflect.ValueOf(a)
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		n := v.Type().Field(j).Name
		switch f.Kind() {
		case reflect.String:
			if f.String() == "" {
				logrus.WithError(errors.New("ENV VARIABLE not found for " + field + n)).Panic(constants.InvalidConfigType)
				return
			}
		case reflect.Struct:
			ValidateConfig(field+n+".", f.Interface())
		}
	}
}

// GetConfig : will give the struct as value so that the actual conf doesn't get tampered
func GetConfig() AppConfig {
	return config
}
