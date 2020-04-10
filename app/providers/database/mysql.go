// Package database will provide the configuration specific to mysql
package database

import (
	"fmt"

	"otm/app/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// MysqlConnectionDSNFormat : DSN for connecting mysql
	MysqlConnectionDSNFormat = "%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
)

type mysql struct {
	database
}

// getDatabasePath will give the connection string for the mysql database connection
func (ms mysql) getDatabasePath(databaseConfig config.DatabaseConfig) string {
	// charset=utf8: uses utf8 character set data format
	// parseTime=true: changes the output type of DATE and DATETIME values to time.Time instead of []byte / strings
	// loc=Local: Sets the location for time.Time values (when using parseTime=true). "Local" sets the system's location
	return fmt.Sprintf(
		MysqlConnectionDSNFormat,
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Protocol,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Name)
}
