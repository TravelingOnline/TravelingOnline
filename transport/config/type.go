package config

type Config struct {
	DB      DBConfig      `json:"db"  yaml:"db"`
	Server  ServerConfig  `json:"server"  yaml:"server"`
	Vehicle VehicleConfig `json:"vehicle"  yaml:"vehicle"`
	Bank    Bank          `json:"bank"  yaml:"bank"`
}

type DBConfig struct {
	Host     string `json:"host"  yaml:"host"`
	Port     uint   `json:"port"  yaml:"port"`
	Database string `json:"database"  yaml:"database"`
	Schema   string `json:"schema"  yaml:"schema"`
	Username string `json:"username"  yaml:"username"`
	Password string `json:"password"  yaml:"password"`
}

type ServerConfig struct {
	HttpPort            uint   `json:"httpPort"  yaml:"httpPort"`
	Secret              string `json:"secret"  yaml:"secret"`
	AuthExpMinute       uint   `json:"authExpMin"  yaml:"authExpMin"`
	AuthRefreshMinute   uint   `json:"authExpRefreshMin"  yaml:"authExpRefreshMin"`
	RateLimitMaxAttempt int    `json:"rate_limit_max_attempt"  yaml:"rate_limit_max_attempt"`
	RatelimitTimePeriod int    `json:"ratelimit_time_period"  yaml:"ratelimit_time_period"`
}

type VehicleConfig struct {
	HttpPort uint   `json:"httpPort"  yaml:"httpPort"`
	Host     string `json:"host"  yaml:"host"`
}

type Bank struct {
	HttpPort uint   `json:"httpPort"  yaml:"httpPort"`
	Host     string `json:"host"  yaml:"host"`
}
