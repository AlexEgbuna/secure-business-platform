package database

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/AlexEgbuna/secure-business-platform/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxConns              int32 = 10
	defaultMinConns              int32 = 2
	defaultMaxConnLifetime             = time.Hour
	defaultMaxConnLifetimeJitter       = 5 * time.Minute
	defaultMaxConnIdleTime             = 30 * time.Minute
	defaultConnectTimeout              = 5 * time.Second
	defaultHealthTimeout               = 5 * time.Second
)

// DB manages PostgreSQL connections for the application.
type DB struct {
	pool *pgxpool.Pool
}

// New creates a PostgreSQL connection pool.
func New(ctx context.Context, cfg config.DatabaseConfig) (*DB, error) {
	if cfg.Host == "" {
		return nil, fmt.Errorf("database host cannot be empty")
	}

	if cfg.Port <= 0 || cfg.Port > 65535 {
		return nil, fmt.Errorf("database port is invalid")
	}

	if cfg.Name == "" {
		return nil, fmt.Errorf("database name cannot be empty")
	}

	if cfg.User == "" {
		return nil, fmt.Errorf("database user cannot be empty")
	}

	if cfg.Password == "" {
		return nil, fmt.Errorf("database password cannot be empty")
	}

	if cfg.SSLMode != "require" && cfg.SSLMode != "disable" {
		return nil, fmt.Errorf("unsupported database SSL mode: %q", cfg.SSLMode)
	}

	connectionURL := url.URL{
		Scheme: "postgres",
		Host:   net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Path:   "/" + cfg.Name,
		User:   url.UserPassword(cfg.User, cfg.Password),
	}

	query := url.Values{}
	query.Set("sslmode", cfg.SSLMode)
	connectionURL.RawQuery = query.Encode()

	poolConfig, err := pgxpool.ParseConfig(connectionURL.String())
	if err != nil {
		return nil, fmt.Errorf("parse database pool configuration")
	}

	poolConfig.MaxConns = defaultMaxConns
	poolConfig.MinConns = defaultMinConns
	poolConfig.MaxConnLifetime = defaultMaxConnLifetime
	poolConfig.MaxConnLifetimeJitter = defaultMaxConnLifetimeJitter
	poolConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	poolConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	if poolConfig.ConnConfig.RuntimeParams == nil {
		poolConfig.ConnConfig.RuntimeParams = make(map[string]string)
	}

	poolConfig.ConnConfig.RuntimeParams["application_name"] =
		"secure-business-platform"

	if cfg.SSLMode == "require" {
		if poolConfig.ConnConfig.TLSConfig == nil {
			return nil, fmt.Errorf("TLS configuration was not created for required database TLS")
		}

		poolConfig.ConnConfig.TLSConfig.MinVersion = tls.VersionTLS12
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create database pool")
	}

	return &DB{pool: pool}, nil
}

// Ping verifies database connectivity.
func (db *DB) Ping(ctx context.Context) error {
	if db == nil || db.pool == nil {
		return fmt.Errorf("database pool is not initialized")
	}

	if err := db.pool.Ping(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}

// Close releases database connections.
func (db *DB) Close() {
	if db == nil || db.pool == nil {
		return
	}

	db.pool.Close()
}

// Pool returns the underlying connection pool.
func (db *DB) Pool() *pgxpool.Pool {
	if db == nil {
		return nil
	}

	return db.pool
}

// HealthCheck verifies database availability with a bounded timeout.
func (db *DB) HealthCheck() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		defaultHealthTimeout,
	)
	defer cancel()

	return db.Ping(ctx)
}
