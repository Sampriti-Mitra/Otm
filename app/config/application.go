package config

type application struct {
	Name       string `toml:"app_name"`
	Mode       string `toml:"app_mode"`
	Env        string `toml:"env"`
	ListenPort int    `toml:"listen_port"`
	ListenIP   string `toml:"listen_ip"`
	LogPath    string `toml:"log_path"`
	PrettyLog  bool   `toml:"pretty_log"`
	InfoLog    bool   `toml:"info_log"`
	Verbose    bool   `toml:"verbose"`
}
