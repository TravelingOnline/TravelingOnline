package config

type Config struct {
	DB     DBConfig     `json:"db"  yaml:"db"`
	Server ServerConfig `json:"server"  yaml:"server"`
}

type DBConfig struct {
	Host      string `json:"host"  yaml:"host"`
	Port      uint   `json:"port"  yaml:"port"`
	QDatabase string `json:"q_database"  yaml:"q_database"`
	SDatabase string `json:"s_database"  yaml:"s_database"`
	Schema    string `json:"schema"  yaml:"schema"`
	Username  string `json:"username"  yaml:"username"`
	Password  string `json:"password"  yaml:"password"`
}

type ServerConfig struct {
	HttpPort            uint   `json:"httpPort"  yaml:"httpPort"`
	Secret              string `json:"secret"  yaml:"secret"`
	AuthExpMinute       uint   `json:"authExpMin"  yaml:"authExpMin"`
	AuthRefreshMinute   uint   `json:"authExpRefreshMin"  yaml:"authExpRefreshMin"`
	RateLimitMaxAttempt int    `json:"rate_limit_max_attempt"  yaml:"rate_limit_max_attempt"`
	RatelimitTimePeriod int    `json:"ratelimit_time_period"  yaml:"ratelimit_time_period"`
}
