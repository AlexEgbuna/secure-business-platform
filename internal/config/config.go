package config

// Config is the root configuration for the Secure Business Platform.
type Config struct {
	Application ApplicationConfig
	Database    DatabaseConfig
	Redis       RedisConfig
	JWT         JWTConfig
	Session     SessionConfig
	Logging     LoggingConfig
	Metrics     MetricsConfig
	GoTLS       GoTLSConfig
	Proxy       ProxyConfig
}

// ApplicationConfig contains general application settings.
type ApplicationConfig struct {
	Name        string
	Environment string
	Host        string
	Port        int
}

// DatabaseConfig contains PostgreSQL configuration.
type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// RedisConfig contains Redis configuration.
type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

// JWTConfig contains JWT configuration.
type JWTConfig struct {
	Secret            string
	Expiration        string
	RefreshExpiration string
}

// SessionConfig contains session configuration.
type SessionConfig struct {
	Secret  string
	Timeout string
}

// LoggingConfig contains logging configuration.
type LoggingConfig struct {
	Level string
}

// MetricsConfig contains metrics configuration.
type MetricsConfig struct {
	Enabled bool
	Port    int
}

// GoTLSConfig contains GoTLS integration settings.
type GoTLSConfig struct {
	Enabled     bool
	ProxyURL    string
	SharedToken string
}

// ProxyConfig contains outbound SOCKS5 proxy configuration.
type ProxyConfig struct {
	Enabled  bool
	Host     string
	Port     int
	Username string
	Password string
}
