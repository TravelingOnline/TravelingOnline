package config

type Config struct {
	AgencyDB DBConfig     `json:"agency_db"  yaml:"agency_db"`
	Logger   LoggerConfig `json:"logger"  yaml:"logger"`
	Server   ServerConfig `json:"server"  yaml:"server"`
}

type ServerConfig struct {
	HttpPort              uint            `json:"httpPort"  yaml:"httpPort"`
	Secret                string          `json:"secret"  yaml:"secret"`
	AuthExpMinute         uint            `json:"authExpMin"  yaml:"authExpMin"`
	AuthRefreshMinute     uint            `json:"authExpRefreshMin"  yaml:"authExpRefreshMin"`
	RateLimitMaxAttempt   int             `json:"rate_limit_max_attempt"  yaml:"rate_limit_max_attempt"`
	RatelimitTimePeriod   int             `json:"ratelimit_time_period"  yaml:"ratelimit_time_period"`
	GRPCPort              int             `json:"grpcPort"`
	Host                  string          `json:"host"`
	ServiceRegistry       ServiceRegistry `json:"service_registry"`
	ServiceHostAddress    string          `json:"service_host_address"`
	ServiceHTTPHealthPath string          `json:"service_http_health_path"`
	ServiceHTTPPrefixPath string          `json:"service_http_prefix_path"`
}
type ServiceRegistry struct {
	Address         string `json:"address"`
	ServiceName     string `json:"service_name"`
	BankServiceName string `json:"bank_service_name"`
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

type RedisConfig struct {
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type MessageBrokerConfig struct {
	Username              string `json:"username"`
	Password              string `json:"password"`
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	CreateWalletQueueName string `json:"CreateWalletQueueName"`
	TransferQueueName     string `json:"TransferQueueName"`
	UserQueueName         string `json:"UserQueueName"`
}

type LoggerConfig struct {
	Level  string `json:"level"  yaml:"level"`
	Output string `json:"output"  yaml:"output"`
	Path   string `json:"path"  yaml:"path"`
}

type EsConfig struct {
	Host     string `json:"host"  yaml:"host"`
	Port     string `json:"port"  yaml:"port"`
	Username string `json:"username"  yaml:"username"`
	Password string `json:"password"  yaml:"password"`
	Index    string `json:"index"  yaml:"index"`
}
