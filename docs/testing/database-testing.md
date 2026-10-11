# Database Testing

## Purpose

Document the verification performed for the PostgreSQL database layer.

## Unit and Package Tests

The database package tests were run using:

```
go test ./internal/database -count=1
```

The command completed successfully.

The full project test suite was run using:

```
go test ./... -count=1
```

The command completed successfully for the application API, configuration, database, and middleware packages.

## PostgreSQL Integration Test

The PostgreSQL integration test is opt-in and runs only when explicitly enabled.

Enable it using:

```
RUN_POSTGRES_INTEGRATION_TESTS=1
```

The test requires valid database environment variables and a reachable PostgreSQL instance.

The following command was run successfully:

```
RUN_POSTGRES_INTEGRATION_TESTS=1 go test ./internal/database -run '^TestPostgreSQLConnection$' -count=1 -v
```

The test verified database connectivity and reported:

```
PostgreSQL integration test passed: user=secure_business_app database=secure_business_platform TLS=true
```

The test was run against the local PostgreSQL instance.

## PostgreSQL TLS Verification

A PostgreSQL client connection was also checked using `psql`.

The observed connection reported TLS 1.3 and the cipher `TLS_AES_256_GCM_SHA384`.

This verifies that the tested connection was encrypted. It does not establish that the local snake-oil certificate is suitable for production or that the server's identity was verified.

## Application Startup and Health Endpoint

The application was started using:

```
go run ./cmd/server
```

The server reported that it was listening on:

```
127.0.0.1:8080
```

The health endpoint was checked using:

```
curl -i http://127.0.0.1:8080/health
```

The endpoint returned HTTP 200 OK with the response:

```
{"status":"ok"}
```

This verifies the tested HTTP health endpoint response. It does not independently establish database health after startup.

## Shutdown Verification

The application was stopped using Ctrl+C.

The application logged receipt of the interrupt signal, confirming that the SIGINT shutdown path was reached during this test.

SIGTERM handling had already been tested successfully in Milestone 2 and was not repeated in Milestone 4.

## Additional Verification

The following checks completed successfully:

```
gofmt -w internal/config internal/database cmd/server
git diff --check
go vet ./...
go build ./...
go mod tidy
go mod verify
go list -m all | grep 'github.com/jackc/pgx'
```

The selected pgx module version was:

```
github.com/jackc/pgx/v5 v5.6.0
```

The Go module verification reported that all modules were verified.

## Issues Encountered and Resolved

### Go Version Compatibility

The initial test implementation used `t.Context()`, which was unavailable in the project's Go 1.22.2 environment.

The test was updated to use `context.Background()`.

### Missing Database Configuration

An initial integration-test attempt failed because the required `DB_NAME` environment variable was missing.

The database environment variables were supplied, and the integration test subsequently passed.

### Application Configuration Validation

An initial application startup attempt was rejected because the JWT secret was missing.

The configuration validation correctly prevented startup with incomplete required configuration. The application subsequently started with the intended local configuration.

## Current Limitations

* Integration testing requires an available PostgreSQL instance and explicit opt-in.
* Local TLS testing does not establish production certificate trust.
* The health endpoint test does not independently verify ongoing database connectivity.
* No production load test or connection-pool benchmark was performed as part of this milestone.
