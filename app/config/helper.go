package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"otm/app/constants"
)

// loadConfigFromFile: load/overwrite config values from given basepath, filename and env
func loadConfigFromFile(basePath string, filename string, env string) {
	path := getFilePath(basePath, filename, env)
	_, err := os.Stat(path)
	if err != nil {
		logrus.WithError(err).Error(constants.ConfigFileNotFound)
	}

	content := readConfigFile(path)

	// inject environment variables
	content = os.ExpandEnv(string(content))

	if _, err := toml.Decode(content, &config); err != nil {
		logrus.WithError(err).Panic(constants.InvalidConfigType)
	}
}

// readConfigFile: config file will be read from the given file and gives the content in string format
func readConfigFile(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		logrus.WithError(err).Error(constants.ConfigFileNotFound)
		return ""
	}

	return string(data)
}

// getFilePath: gives the file path based on the environment provided
// file path will be relative to the application and determined by basePath
func getFilePath(basePath string, fileName string, env string) string {
	if env != "" {
		fileName = fmt.Sprintf(fileName, env)
	}

	path := fmt.Sprintf(FilePath, basePath, fileName)

	return path
}
