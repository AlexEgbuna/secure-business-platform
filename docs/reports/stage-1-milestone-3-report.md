# Stage 1 — Milestone 3 Engineering Report

## Objective

Establish application-level HTTP routing for the Secure Business Platform while keeping application routing 
separate from the HTTP server lifecycle and the GoTLS edge/security layer.

## Scope

- Application router under `internal/api`.
- Health endpoint.
- Version endpoint.
- HTTP method restrictions.
- Middleware abstraction under `internal/middleware`.
- Route-group registration architecture.
- Integration with `cmd/server`.
- Router and middleware testing.
- Application-routing security considerations.

Authentication, authorization, database access, business modules, WAF, TLS termination, reverse proxying, and GoTLS implementation are outside this milestone.

## Application Router Architecture

Application route ownership is located in:

    internal/api/

The process entry point remains responsible for:

- configuration;
- HTTP server construction;
- startup;
- signal handling;
- graceful shutdown.

`internal/api.Router` owns the application `http.ServeMux`.

The router uses Go's standard-library `http.ServeMux` with method-aware route patterns.

## Route Registration

Current application routes are:

    GET /health
    GET /version

Unknown routes return `404 Not Found`.

Unsupported methods return `405 Method Not Allowed`.

Future application modules can register routes through:

    type RouteGroup interface {
        RegisterRoutes(*http.ServeMux)
    }

This provides a controlled mechanism for future application modules to register their routes without moving route ownership into `cmd/server`.

## Health Endpoint

The health endpoint returns:

    {"status":"ok"}

Only minimal health information is exposed.

The endpoint does not expose application configuration, credentials, or infrastructure details.

## Version Endpoint

The version endpoint currently returns:

    {"version":"dev"}

The version is represented by an application variable so a future build/release process can inject the production version.

## Middleware Architecture

Middleware is represented by:

    type Middleware func(http.Handler) http.Handler

Middleware is composed through:

    func Chain(handler http.Handler, middlewares ...Middleware) http.Handler

The first middleware supplied is the outermost middleware and therefore executes first.

Middleware ordering is deterministic and tested.

## HTTP Method Restrictions

The router uses method-aware route patterns.

Current routes explicitly accept `GET`.

Unsupported methods return `405 Method Not Allowed`.

Unknown paths return `404 Not Found`.

## Security Considerations

The current router exposes only the routes explicitly registered by the application.

The health response does not expose configuration or secrets.

The version endpoint exposes only the current development version identifier.

The router does not provide authentication or authorization. Those controls belong to later identity and application-security milestones.

Middleware ordering is deterministic and tested because future authentication, authorization, audit, and 
security middleware will depend on predictable execution order.

## GoTLS Boundary

The application router is not a reverse proxy, WAF, TLS terminator, or edge security engine.

GoTLS remains responsible for its defined edge/network security responsibilities.

The Business Platform remains responsible for application routing and application security decisions.

The future integration boundary will include TLS encryption from GoTLS to the Business application.

## Verification Performed

Milestone verification established successful:

- source formatting;
- unit testing;
- static analysis;
- project compilation;
- server compilation;
- HTTP functional behavior;
- method restrictions;
- unknown-route handling;
- middleware execution;
- route-group registration.

## Test Results

The router tests cover:

- health endpoint;
- version endpoint;
- unknown routes;
- HTTP method restrictions;
- middleware execution;
- route-group registration.

The middleware tests verify deterministic execution order.

Functional testing confirmed:

    GET /health          200
    GET /version         200
    GET /does-not-exist  404
    POST /health        405
    POST /version       405
    GET /random-route   404

## Git Verification

The Milestone 3 implementation is prepared for the milestone commit:

    feat: implement application router

## Security Review

The application routing layer does not claim to provide authentication or authorization.

No application secrets are returned by the health or version endpoints.

Unknown routes are rejected.

Unsupported HTTP methods are rejected.

The application router remains separated from the GoTLS edge-security boundary.

## Engineering Decisions

- Standard-library `http.ServeMux` is used.
- Application routes are separated from `cmd/server`.
- Middleware has its own package.
- Future application modules use `RouteGroup`.
- Health output is intentionally minimal.
- Production version injection is deferred to the build/release process.
- GoTLS edge responsibilities are not duplicated inside the application router.

## Lessons Learned

Application lifecycle and application behavior should remain separated.

Routing should have explicit ownership.

Middleware ordering must be deterministic.

Security controls should be implemented explicitly rather than assumed to exist because an application router exists.

Application routing must remain separate from edge/network security responsibilities.

## Milestone Completion

Milestone 3 documentation records the implemented application-routing architecture, security considerations, 
testing results, engineering decisions, and verification results.
