package config

import "time"

// authUser: All the users that are going to interact with Scrooge
type authUser struct {
	//CPS           AuthUserConfig `toml:"cps"`
	//ROUTINGENGINE AuthUserConfig `toml:"routingengine"`
	//ADMINAPI      AuthUserConfig `toml:"adminapi"`
	//LUMBERJACK    AuthUserConfig `toml:"lumberjack"`
}

// AuthUserConfig : Stores credentials config of a given authUser
type AuthUserConfig struct {
	UserName string `toml:"username"`
	Password string `toml:"password"`
}

type Auth struct {
	Key    string `toml:"key"`
	Secret string `toml:"secret"`
}

type Endpoint struct {
	URL     string            `toml:"url"`
	Method  string            `toml:"method"`
	Timeout time.Duration     `toml:"timeout"`
	Auth    Auth              `toml:"auth"`
	Headers map[string]string `toml:"headers"`
}
