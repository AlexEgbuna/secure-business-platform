# Stage 1 — Milestone 4 Report: Database Layer

## Objective

Implement the initial PostgreSQL database layer for the Secure Business Platform and integrate database initialization and
connectivity verification into application startup.

## Scope

This milestone covers:

* PostgreSQL connectivity;
* database configuration;
* connection-pool initialization;
* startup connectivity verification;
* application database lifecycle integration;
* database integration testing;
* applicable architecture, security, testing, and operational documentation.

Database migrations, domain models, authentication changes, business modules, and production deployment work remain outside this milestone.

## Implementation

The database layer was added under:

```
internal/database/
```

The implementation uses `github.com/jackc/pgx/v5` and `pgxpool`.

Database settings are integrated with the existing application configuration. The application initializes the database pool and explicitly
checks connectivity using `Ping` with a bounded context before proceeding with server startup.

The application lifecycle includes closing the database pool during shutdown.

## Dependency Compatibility

The environment uses Go 1.22.2.

The initially considered pgx version required a newer Go version than the project's environment. The implementation therefore uses:

```
github.com/jackc/pgx/v5 v5.6.0
```

This version is compatible with the project's Go version.

Module tidying and module verification completed successfully.

## PostgreSQL Environment

The local environment uses PostgreSQL 16.

The application connects using the dedicated `secure_business_app` database role. The role was checked and does not have superuser, role-creation,
database-creation, replication, or `BYPASSRLS` privileges.

The PostgreSQL service was available, and application connectivity was verified.

## Connection Pool

Initial pool settings include a maximum of 10 connections, a minimum of 2 connections, a five-second connection timeout, a one-hour maximum
connection lifetime, five minutes of connection lifetime jitter, and a 30-minute maximum idle time.

These values establish bounded initial behavior. They have not been selected through production load testing.

## Transport Security

The local application-to-PostgreSQL connection was verified as TLS-encrypted.

The observed connection negotiated TLS 1.3 using `TLS_AES_256_GCM_SHA384`.

The local installation uses a snake-oil certificate, and the tested `sslmode=require` configuration does not verify the server's certificate identity.

Trusted certificate configuration and hostname verification remain necessary before production use.

## Testing and Verification

The following checks completed successfully:

* Database package tests.
* Full Go test suite using `go test ./... -count=1`.
* PostgreSQL integration test with explicit opt-in.
* `go vet ./...`.
* `go build ./...`.
* `gofmt` on the affected Go packages.
* `git diff --check`.
* `go mod tidy`.
* `go mod verify`.

The integration test confirmed the expected application database user, database name, and TLS-enabled connection.

The application started successfully with valid local configuration. The `/health` endpoint returned HTTP 200 OK with `{"status":"ok"}`.

Ctrl+C was used to exercise the SIGINT shutdown path. SIGTERM handling was successfully tested in Milestone 2 and was not repeated in this milestone.

## Issues Encountered and Resolved

### Go API Compatibility

The initial test code used `t.Context()`, which is not available in Go 1.22.2. It was replaced with `context.Background()`.

### Missing Environment Variables

An initial integration-test run failed because required database configuration was missing. The environment was configured, and the integration test
subsequently passed.

### Required Application Configuration

An initial application startup attempt was rejected because the required JWT secret was missing. This demonstrated that configuration validation prevented
startup with incomplete required configuration. The application subsequently started with the intended local configuration.

### Dependency Selection

The selected pgx version was adjusted to maintain compatibility with Go 1.22.2.

## Security Considerations

* Database access uses a dedicated application role.
* Database credentials are supplied through configuration and must not be committed.
* Connection-pool limits provide bounded initial resource use.
* TLS encryption was verified for the local database connection.
* The local certificate and `sslmode=require` setting do not provide production-grade server identity verification.
* Pool settings have not been load-tested.

## Deferred Work

The following work is intentionally deferred:

* Database migration system.
* Domain models and repositories.
* Authentication and authorization implementation.
* Business modules.
* Production TLS certificate and CA configuration.
* Evidence-based connection-pool tuning.
* Production secret management.

## Outcome

The initial PostgreSQL database layer is implemented and the recorded package, integration, build, vet, dependency, and local HTTP checks completed successfully.

The milestone's production limitations are recorded explicitly. Completion of this milestone does not imply that production database certificate verification,
performance tuning, or later application data-access features have been implemented.
