package config

import (
	"fmt"
	"os"
	"strconv"
)

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	cfg := &Config{}

	// --------------------------------------------------
	// Application
	// --------------------------------------------------

	cfg.Application.Name = getEnv("APP_NAME", "Secure Business Platform")
	cfg.Application.Environment = getEnv("APP_ENV", "development")
	cfg.Application.Host = getEnv("APP_HOST", "127.0.0.1")

	appPort, err := getEnvAsInt("APP_PORT", 8080)
	if err != nil {
		return nil, fmt.Errorf("invalid APP_PORT: %w", err)
	}
	cfg.Application.Port = appPort

	// --------------------------------------------------
	// Database
	// --------------------------------------------------

	cfg.Database.Host = getEnv("DB_HOST", "127.0.0.1")

	dbPort, err := getEnvAsInt("DB_PORT", 5432)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}
	cfg.Database.Port = dbPort

	cfg.Database.Name = getEnv("DB_NAME", "")
	cfg.Database.User = getEnv("DB_USER", "")
	cfg.Database.Password = getEnv("DB_PASSWORD", "")

	// --------------------------------------------------
	// Redis
	// --------------------------------------------------

	cfg.Redis.Host = getEnv("REDIS_HOST", "127.0.0.1")

	redisPort, err := getEnvAsInt("REDIS_PORT", 6379)
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_PORT: %w", err)
	}
	cfg.Redis.Port = redisPort

	cfg.Redis.Password = getEnv("REDIS_PASSWORD", "")

	// --------------------------------------------------
	// JWT
	// --------------------------------------------------

	cfg.JWT.Secret = getEnv("JWT_SECRET", "")
	cfg.JWT.Expiration = getEnv("JWT_EXPIRATION", "15m")
	cfg.JWT.RefreshExpiration = getEnv("JWT_REFRESH_EXPIRATION", "168h")

	// --------------------------------------------------
	// Session
	// --------------------------------------------------

	cfg.Session.Secret = getEnv("SESSION_SECRET", "")
	cfg.Session.Timeout = getEnv("SESSION_TIMEOUT", "24h")

	// --------------------------------------------------
	// Logging
	// --------------------------------------------------

	cfg.Logging.Level = getEnv("LOG_LEVEL", "info")

	// --------------------------------------------------
	// Metrics
	// --------------------------------------------------

	cfg.Metrics.Enabled = getEnvAsBool("METRICS_ENABLED", true)

	metricsPort, err := getEnvAsInt("METRICS_PORT", 9090)
	if err != nil {
		return nil, fmt.Errorf("invalid METRICS_PORT: %w", err)
	}
	cfg.Metrics.Port = metricsPort

	// --------------------------------------------------
	// GoTLS
	// --------------------------------------------------

	cfg.GoTLS.Enabled = getEnvAsBool("GOTLS_ENABLED", true)
	cfg.GoTLS.ProxyURL = getEnv("GOTLS_PROXY_URL", "")
	cfg.GoTLS.SharedToken = getEnv("GOTLS_SHARED_TOKEN", "")

	// --------------------------------------------------
	// SOCKS5 Proxy
	// --------------------------------------------------

	cfg.Proxy.Enabled = getEnvAsBool("SOCKS5_ENABLED", false)
	cfg.Proxy.Host = getEnv("SOCKS5_HOST", "")

	proxyPort, err := getEnvAsInt("SOCKS5_PORT", 1080)
	if err != nil {
		return nil, fmt.Errorf("invalid SOCKS5_PORT: %w", err)
	}
	cfg.Proxy.Port = proxyPort

	cfg.Proxy.Username = getEnv("SOCKS5_USERNAME", "")
	cfg.Proxy.Password = getEnv("SOCKS5_PASSWORD", "")

	return cfg, nil
}

// getEnv returns an environment variable or a default value.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// getEnvAsInt returns an integer environment variable.
func getEnvAsInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)

	if value == "" {
		return fallback, nil
	}

	return strconv.Atoi(value)
}

// getEnvAsBool returns a boolean environment variable.
func getEnvAsBool(key string, fallback bool) bool {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return parsed
}
