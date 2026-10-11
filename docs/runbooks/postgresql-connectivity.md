# PostgreSQL Connectivity Runbook

## Purpose

Provide operational steps for checking PostgreSQL availability and diagnosing local database connectivity problems in the Secure Business Platform.

## Environment

The current local development environment uses:

* PostgreSQL 16;
* database `secure_business_platform`;
* application role `secure_business_app`;
* PostgreSQL TCP port `5432`;
* application HTTP address `127.0.0.1:8080`.

These values describe the tested local environment and must be adjusted for other deployments.

## 1. Check PostgreSQL Service Status

Check the PostgreSQL service:

```
sudo systemctl status postgresql
```

Check the cluster:

```
pg_lsclusters
```

Confirm that the expected PostgreSQL cluster is online and listening on the expected port.

## 2. Check Required Environment Variables

The application requires the database connection settings:

```
DB_HOST
DB_PORT
DB_NAME
DB_USER
DB_PASSWORD
DB_SSL_MODE
```

Check whether each variable is set without printing its value:

```
for name in DB_HOST DB_PORT DB_NAME DB_USER DB_PASSWORD DB_SSL_MODE; do
  if [ -n "${!name:-}" ]; then
    printf '%s is set\n' "$name"
  else
    printf '%s is missing\n' "$name"
  fi
done
```

Environment variables set in one shell are not automatically available in a separate terminal or unrelated process. Ensure that the application and test process receive the intended configuration.

Never paste real passwords or secret values into logs, documentation, or chat.

## 3. Test PostgreSQL Connectivity

Use `psql` to check connectivity:

```
psql "host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} sslmode=${DB_SSL_MODE}" -W
```

The `-W` option prompts for the password rather than requiring it to be included in the command text.

After connecting, inspect the current database identity:

```
SELECT current_user, current_database();
```

Check whether the session uses TLS:

```
\conninfo
```

Confirm that the connected role and database are the expected application values.

## 4. Run the PostgreSQL Integration Test

The database integration test is opt-in.

First ensure the required database environment variables are set. Then run:

```
RUN_POSTGRES_INTEGRATION_TESTS=1 go test ./internal/database -run '^TestPostgreSQLConnection$' -count=1 -v
```

The test requires a reachable PostgreSQL instance and valid credentials.

A missing environment variable, authentication failure, incorrect database name, or unavailable server may cause the test to fail.

## 5. Start the Application

With the required application and database configuration available in the current environment, run:

```
go run ./cmd/server
```

The application should report that the HTTP server is listening on its configured address.

If startup fails, read the error carefully. Configuration validation may reject missing required settings before the application attempts to serve requests.

Do not resolve configuration errors by weakening validation or printing secret values.

## 6. Check the Health Endpoint

In another terminal, run:

```
curl -i http://127.0.0.1:8080/health
```

The expected response for the tested local configuration is HTTP 200 OK with:

```
{"status":"ok"}
```

This endpoint confirms the tested HTTP health response. It does not independently prove that PostgreSQL remains reachable after startup.

## 7. Troubleshooting

### PostgreSQL Is Not Running

Inspect the system service and cluster status. Resolve the service or cluster issue before repeating the connection test.

### Required Environment Variable Is Missing

Set the required variable in the environment used by the application or integration test. Avoid printing its value when checking configuration.

### Authentication Fails

Confirm the database role, password source, host-based authentication rules, and intended authentication method. Do not switch to a more privileged database account merely to bypass an authentication problem.

### Database or Host Cannot Be Reached

Check the PostgreSQL service, listening address, port, database name, and network connectivity.

### TLS Connection Fails

Inspect the configured SSL mode and PostgreSQL TLS settings. Do not disable TLS verification as a permanent workaround.

The current local `sslmode=require` setting encrypts the connection but does not verify the server's certificate identity. Production deployments should use a trusted certificate and appropriate CA and hostname verification, such as `sslmode=verify-full`.

### Application Startup Is Rejected

Check the configuration-validation error and confirm that all required settings are present. Do not weaken required configuration validation to make startup succeed.

## 8. Security Reminders

* Do not commit real credentials.
* Do not include passwords in command history or logs.
* Use the dedicated application database role.
* Do not use a PostgreSQL superuser for normal application connections.
* Do not treat TLS encryption as proof of server identity verification.
* Do not weaken authentication or TLS verification to hide connectivity problems.
* Reassess database privileges, certificate configuration, and connection-pool settings before production deployment.
