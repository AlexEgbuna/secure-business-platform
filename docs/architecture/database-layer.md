# Database Layer Architecture

## Purpose

Define the database access boundary for the Secure Business Platform.

## Ownership

Database connectivity and connection-pool management are owned by:

```
internal/database/
```

`cmd/server/main.go` remains responsible for application composition and server lifecycle. It initializes the database layer and ensures the pool is closed during application shutdown.

## Database

The application uses PostgreSQL as its relational database.

Current local development configuration:

* Database: PostgreSQL
* Database driver and connection pool: `github.com/jackc/pgx/v5`
* Connection pool: `pgxpool`
* Database configuration: `internal/config/`
* Database access implementation: `internal/database/`

The selected pgx version is compatible with the project's Go 1.22.2 environment.

## Configuration

Database connection settings are supplied through application configuration and environment variables.

Current settings include:

* `DB_HOST`
* `DB_PORT`
* `DB_NAME`
* `DB_USER`
* `DB_PASSWORD`
* `DB_SSL_MODE`

The `.env.example` file and `config.example.json` provide example configuration. They must not contain real credentials.

## Connection Pool

The database layer uses `pgxpool` to manage PostgreSQL connections.

The initial pool configuration includes:

* Maximum connections: 10
* Minimum connections: 2
* Maximum connection lifetime: 1 hour
* Connection lifetime jitter: 5 minutes
* Maximum connection idle time: 30 minutes
* Connection timeout: 5 seconds

These are initial development settings, not values established through production load testing.

The application explicitly pings PostgreSQL during startup so that database connectivity is checked before the application proceeds.

## Lifecycle

The database lifecycle is integrated into application startup and shutdown.

The startup process:

1. Loads and validates application configuration.
2. Creates the database connection pool.
3. Checks database connectivity using `Ping`.
4. Starts the HTTP server after successful initialization.

The database pool is closed when the application shuts down.

## Transport Security

The local PostgreSQL connection was verified using TLS.

The current local configuration supports TLS 1.2 or later and was observed negotiating TLS 1.3 during verification.

The local PostgreSQL installation uses a snake-oil certificate. TLS encryption was verified, but the current `sslmode=require` configuration does not verify the server's certificate identity.

Production deployment must use an appropriately trusted certificate and certificate verification, such as `sslmode=verify-full`, with correct hostname and CA configuration.

## Health Checks

The existing `/health` endpoint returns the application's health response.

A successful HTTP health response alone does not prove that PostgreSQL is reachable. Database connectivity is checked during application startup. Any future database-aware health-check behavior must be designed explicitly rather than assumed.

## GoTLS Boundary

The database layer is not:

* an application router;
* an edge reverse proxy;
* a WAF;
* a TLS termination service;
* a replacement for GoTLS.

The Business Platform owns application configuration and database access. GoTLS remains the designated edge and network traffic-security layer.

## Current Scope

Implemented:

* PostgreSQL connection-pool initialization;
* database configuration integration;
* startup connectivity verification;
* database pool lifecycle integration;
* opt-in PostgreSQL integration testing.

Deferred:

* database migrations;
* domain models and repositories;
* authentication and authorization changes;
* business modules;
* production certificate deployment;
* database performance tuning and index design.
