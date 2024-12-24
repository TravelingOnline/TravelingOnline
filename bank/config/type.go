package config

type Config struct {
	BANK_DB DBConfig     `json:"bank_db"  yaml:"bank_db"`
	Logger  LoggerConfig `json:"logger"  yaml:"logger"`
	Server  ServerConfig `json:"server"  yaml:"server"`
}

type ServerConfig struct {
	HttpPort            int `json:"httpPort"  yaml:"httpPort"`
	GRPCPort            int `json:"grpcPort"`
	RateLimitMaxAttempt int `json:"rate_limit_max_attempt"  yaml:"rate_limit_max_attempt"`
	RatelimitTimePeriod int `json:"ratelimit_time_period"  yaml:"ratelimit_time_period"`

	Host                  string          `json:"host"`
	ServiceRegistry       ServiceRegistry `json:"service_registry"`
	ServiceHostAddress    string          `json:"service_host_address"`
	ServiceHTTPHealthPath string          `json:"service_http_health_path"`
	ServiceHTTPPrefixPath string          `json:"service_http_prefix_path"`
}

type ServiceRegistry struct {
	Address         string `json:"address"`
	ServiceName     string `json:"service_name"`
	AuthServiceName string `json:"auth_service_name"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type LoggerConfig struct {
	Level  string `json:"level"  yaml:"level"`
	Output string `json:"output"  yaml:"output"`
	Path   string `json:"path"  yaml:"path"`
}
