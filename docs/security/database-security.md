# Database Security

## Purpose

Document the security controls and current limitations of the Secure Business Platform's PostgreSQL integration.

## Database Access

The application connects to PostgreSQL using the dedicated role:

```
secure_business_app
```

The role was verified not to have superuser, role-creation, database-creation, replication, or `BYPASSRLS` privileges.

The application must use a dedicated database identity rather than PostgreSQL's administrative account.

Database privileges must remain limited to the operations required by the application as its functionality grows.

## Credential Management

Database credentials are provided through environment-based configuration.

The following rules apply:

* Real database passwords must not be committed to Git.
* Example configuration must use placeholders rather than real credentials.
* Credentials must not be printed in logs or error messages.
* Secrets must not be included in documentation, test output, or commit messages.
* Local shell environment variables are temporary and are available only to processes that inherit them.

Production secret storage and rotation will be addressed as part of the appropriate deployment and security stages.

## Authentication

The PostgreSQL host authentication configuration uses SCRAM authentication for password-based TCP connections.

Local operating-system connections may use PostgreSQL peer authentication, depending on the connection method and configured authentication rules.

Application connections must use the intended application database role and authentication method.

## Transport Encryption

PostgreSQL TLS encryption was verified from the application environment.

The observed connection negotiated:

* Protocol: TLS 1.3
* Cipher: TLS_AES_256_GCM_SHA384

The local PostgreSQL configuration requires TLS 1.2 or later.

The local installation currently uses a snake-oil certificate. The connection was encrypted, but the application's `sslmode=require` configuration does not verify the server's identity.

For production, configure a trusted server certificate and CA, ensure the hostname matches the certificate, and use certificate-verifying connection settings such as `sslmode=verify-full`.

Encryption and server identity verification are separate security properties and must not be treated as interchangeable.

## Connection Pool Limits

The database connection pool has initial limits to prevent unbounded connection creation.

Current settings include:

* Maximum connections: 10
* Minimum connections: 2
* Connection timeout: 5 seconds

These settings have not been validated through production load testing and may require evidence-based adjustment.

## Error Handling and Logging

Database errors must be handled without exposing passwords, connection strings containing credentials, or other secrets.

Connection failures should provide enough operational information to diagnose the issue without disclosing sensitive configuration.

## Current Limitations

* The local certificate is not suitable as a production trust configuration.
* `sslmode=require` does not verify the PostgreSQL server's certificate identity.
* Connection-pool values are initial settings rather than load-tested limits.
* Database migrations and application data-access features are outside this milestone.
* Production secret storage and rotation are not implemented as part of this milestone.

## Security Follow-Up

Before production deployment:

* Configure a trusted PostgreSQL certificate and CA.
* Verify the server identity using `sslmode=verify-full`.
* Review database privileges against the application's actual required operations.
* Evaluate connection-pool settings using measured workload.
* Establish appropriate production secret-management procedures.
