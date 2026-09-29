package api

import (
	"net/http"

	"github.com/AlexEgbuna/secure-business-platform/internal/middleware"
)

// RouteGroup represents an application module that can register
// its HTTP routes with the application router.
//
// Future modules such as authentication, users, organizations,
// roles, dashboard, and audit can implement this interface.
type RouteGroup interface {
	RegisterRoutes(*http.ServeMux)
}

// Router owns application-level HTTP route registration.
//
// It is deliberately separate from the HTTP server lifecycle and
// from the GoTLS edge/security layer.
type Router struct {
	mux         *http.ServeMux
	middlewares []middleware.Middleware
}

// NewRouter creates the application router and registers the routes
// owned by the current milestone.
func NewRouter(middlewares ...middleware.Middleware) *Router {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /version", versionHandler)

	return &Router{
		mux:         mux,
		middlewares: middlewares,
	}
}

// RegisterGroup allows future application modules to register their
// own routes without moving route ownership into cmd/server.
func (r *Router) RegisterGroup(group RouteGroup) {
	group.RegisterRoutes(r.mux)
}

// Handler returns the complete application HTTP handler with the
// configured middleware chain applied.
func (r *Router) Handler() http.Handler {
	return middleware.Chain(r.mux, r.middlewares...)
}
