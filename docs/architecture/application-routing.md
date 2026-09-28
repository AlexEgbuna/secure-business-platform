# Application Routing Architecture

## Purpose

Define the application-level HTTP routing boundary for the Secure Business Platform.

## Ownership

Application HTTP routes are owned by:

    internal/api/

`cmd/server/main.go` remains responsible for application composition and HTTP server lifecycle.

## Router

The application uses Go's standard-library `http.ServeMux`.

Current routes:

    GET /health
    GET /version

Method-aware route patterns provide explicit HTTP method restrictions.

## Middleware

Middleware is implemented under:

    internal/middleware/

The middleware abstraction is:

    type Middleware func(http.Handler) http.Handler

Middleware is composed through:

    func Chain(handler http.Handler, middlewares ...Middleware) http.Handler

The first middleware supplied is the outermost middleware.

## Route Groups

Future application modules can implement:

    type RouteGroup interface {
        RegisterRoutes(*http.ServeMux)
    }

This allows application modules to register routes without moving route ownership into `cmd/server`.

## GoTLS Boundary

The application router is not:

- a reverse proxy;
- a WAF;
- a TLS terminator;
- an edge traffic-security engine;
- a replacement for GoTLS.

GoTLS remains the designated edge/network security layer.

The Business Platform owns application routing and application security decisions.

The future GoTLS-to-Business-Platform integration boundary will include TLS encryption between the two layers.

## Current Scope

Implemented:

- application router;
- health route;
- version route;
- middleware composition;
- route-group registration.

Deferred:

- authentication;
- authorization;
- database integration;
- business modules;
- GoTLS integration.
