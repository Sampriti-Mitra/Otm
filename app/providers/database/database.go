package database

import (
	"otm/app/config"
	"otm/app/interfaces"
)

type database interface {
	getDatabasePath() string
}

// Client holds the database connection
var (
	client *dbProvider
)

// Initialize will initialize the connection to the dialect
func Initialize() {
	if client == nil {
		client = getDbProvider()
		client.connect()
	}
}

// GetClient will give the data base client for the mode set in context
func GetClient() interfaces.IDbProvider {
	return client
}

// GetDatabasePath will give the connection string for the mysql database connection
// More dialect connection string to be handled here in case needed
func GetDatabasePath(dialect string, databaseConfig config.DatabaseConfig) string {
	switch dialect {
	case "mysql":
		return new(mysql).getDatabasePath(databaseConfig)
	default:
		return new(mysql).getDatabasePath(databaseConfig)
	}
}

// Creates a new database provider for the dialect
func getDbProvider() *dbProvider {
	databaseConfig := config.GetConfig().Database

	dbClient := new(dbProvider)
	dbClient.dialect = databaseConfig.Dialect
	dbClient.maxIdleConnections = databaseConfig.MaxIdleConnections
	dbClient.maxOpenConnections = databaseConfig.MaxOpenConnections
	dbClient.connMaxLifetime = databaseConfig.ConnectionMaxLifetime
	dbClient.path = GetDatabasePath(dbClient.dialect, databaseConfig)

	return dbClient
}
