# Application Routing Security

## Route Exposure

The application currently exposes:

    GET /health
    GET /version

Unknown paths return:

    404 Not Found

## Method Restrictions

Current routes explicitly accept `GET`.

Unsupported methods return:

    405 Method Not Allowed

## Information Disclosure

The health endpoint exposes only:

    {"status":"ok"}

The version endpoint currently exposes:

    {"version":"dev"}

The router does not expose:

- database credentials;
- Redis credentials;
- JWT secrets;
- session secrets;
- proxy credentials;
- GoTLS shared tokens.

## Middleware Security

Middleware ordering is deterministic and tested.

This provides the foundation required for future authentication, authorization, audit, and security middleware.

## Authentication and Authorization

Authentication and authorization are not implemented by the application router.

Future protected routes must explicitly implement the appropriate access-control requirements.

## GoTLS Separation

The application router does not implement:

- TLS termination;
- reverse proxying;
- WAF functionality;
- edge traffic filtering;
- GoTLS stateful security.

These remain outside the application router boundary.

## Future Security Work

Future milestones address:

- authentication;
- authorization;
- session security;
- identity management;
- audit logging;
- database security;
- application security;
- GoTLS/application encrypted transport.
