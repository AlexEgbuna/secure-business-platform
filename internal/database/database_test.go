package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/AlexEgbuna/secure-business-platform/internal/config"
)

func TestNewRejectsUnsupportedSSLMode(t *testing.T) {
	cfg := config.DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		Name:     "secure_business_platform",
		User:     "secure_business_app",
		Password: "test-password",
		SSLMode:  "invalid",
	}

	_, err := New(context.Background(), cfg)
	if err == nil {
		t.Fatal("expected unsupported SSL mode to return an error")
	}
}

func TestPostgreSQLConnection(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping PostgreSQL integration test in short mode")
	}

	if os.Getenv("RUN_POSTGRES_INTEGRATION_TESTS") != "1" {
		t.Skip("set RUN_POSTGRES_INTEGRATION_TESTS=1 to run PostgreSQL integration test")
	}

	cfg, err := databaseConfigFromEnvironment()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("create database pool: %v", err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		t.Fatalf("ping PostgreSQL: %v", err)
	}

	var currentUser string
	var currentDatabase string

	err = db.Pool().QueryRow(
		ctx,
		"SELECT current_user, current_database()",
	).Scan(&currentUser, &currentDatabase)
	if err != nil {
		t.Fatalf("query PostgreSQL identity: %v", err)
	}

	if currentUser != cfg.User {
		t.Fatalf("unexpected database user: got %q", currentUser)
	}

	if currentDatabase != cfg.Name {
		t.Fatalf("unexpected database name: got %q", currentDatabase)
	}

	var sslEnabled bool
	var tlsVersion string
	var cipher string

	err = db.Pool().QueryRow(
		ctx,
		`SELECT ssl, version, cipher
		 FROM pg_stat_ssl
		 WHERE pid = pg_backend_pid()`,
	).Scan(&sslEnabled, &tlsVersion, &cipher)
	if err != nil {
		t.Fatalf("query PostgreSQL TLS status: %v", err)
	}

	if cfg.SSLMode == "require" && !sslEnabled {
		t.Fatal("expected PostgreSQL connection to use TLS")
	}

	if cfg.SSLMode == "require" && tlsVersion == "" {
		t.Fatal("expected PostgreSQL TLS version to be reported")
	}

	if cfg.SSLMode == "require" && cipher == "" {
		t.Fatal("expected PostgreSQL TLS cipher to be reported")
	}

	t.Logf(
		"PostgreSQL integration test passed: user=%s database=%s TLS=%t",
		currentUser,
		currentDatabase,
		sslEnabled,
	)
}

func databaseConfigFromEnvironment() (config.DatabaseConfig, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	portValue := os.Getenv("DB_PORT")
	if portValue == "" {
		portValue = "5432"
	}

	port, err := strconv.Atoi(portValue)
	if err != nil || port <= 0 || port > 65535 {
		return config.DatabaseConfig{}, fmt.Errorf("DB_PORT must be a valid TCP port")
	}

	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("DB_SSL_MODE")
	if sslMode == "" {
		sslMode = "require"
	}

	if name == "" {
		return config.DatabaseConfig{}, fmt.Errorf("DB_NAME is required for PostgreSQL integration testing")
	}

	if user == "" {
		return config.DatabaseConfig{}, fmt.Errorf("DB_USER is required for PostgreSQL integration testing")
	}

	if password == "" {
		return config.DatabaseConfig{}, fmt.Errorf("DB_PASSWORD is required for PostgreSQL integration testing")
	}

	return config.DatabaseConfig{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     user,
		Password: password,
		SSLMode:  sslMode,
	}, nil
}
