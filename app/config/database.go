package config

import "time"

type DatabaseConfig struct {
	Dialect               string        `toml:"dialect"`
	Protocol              string        `toml:"protocol"`
	Host                  string        `toml:"host"`
	Port                  int           `toml:"port"`
	Username              string        `toml:"username"`
	Password              string        `toml:"password"`
	Name                  string        `toml:"name"`
	MaxOpenConnections    int           `toml:"max_open_connections"`
	MaxIdleConnections    int           `toml:"max_idle_connections"`
	ConnectionMaxLifetime time.Duration `toml:"conn_max_lifetime"`
}
