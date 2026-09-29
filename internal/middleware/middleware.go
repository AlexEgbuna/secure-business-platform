package middleware

import "net/http"

// Middleware represents an HTTP middleware component.
//
// A middleware receives the next HTTP handler in the request pipeline
// and returns a new handler that wraps it.
type Middleware func(http.Handler) http.Handler

// Chain composes middleware around an HTTP handler.
//
// Middleware is applied in the order provided:
// the first middleware is the outermost layer and therefore executes first.
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
