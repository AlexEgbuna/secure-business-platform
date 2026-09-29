package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterHealth(t *testing.T) {
	router := NewRouter()

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	router.Handler().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if contentType := response.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", contentType)
	}

	expectedBody := "{\"status\":\"ok\"}\n"
	if response.Body.String() != expectedBody {
		t.Fatalf("expected body %q, got %q", expectedBody, response.Body.String())
	}
}

func TestRouterVersion(t *testing.T) {
	router := NewRouter()

	request := httptest.NewRequest(http.MethodGet, "/version", nil)
	response := httptest.NewRecorder()

	router.Handler().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if contentType := response.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", contentType)
	}

	expectedBody := "{\"version\":\"dev\"}\n"
	if response.Body.String() != expectedBody {
		t.Fatalf("expected body %q, got %q", expectedBody, response.Body.String())
	}
}

func TestRouterUnknownRoute(t *testing.T) {
	router := NewRouter()

	request := httptest.NewRequest(http.MethodGet, "/does-not-exist", nil)
	response := httptest.NewRecorder()

	router.Handler().ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, response.Code)
	}
}

func TestRouterHealthMethodRestriction(t *testing.T) {
	router := NewRouter()

	request := httptest.NewRequest(http.MethodPost, "/health", nil)
	response := httptest.NewRecorder()

	router.Handler().ServeHTTP(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, response.Code)
	}
}

func TestRouterMiddlewareExecution(t *testing.T) {
	executed := false

	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			executed = true
			next.ServeHTTP(w, r)
		})
	}

	router := NewRouter(testMiddleware)

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	router.Handler().ServeHTTP(response, request)

	if !executed {
		t.Fatal("expected middleware to execute")
	}

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}
}

type testRouteGroup struct{}

func (testRouteGroup) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("test"))
	})
}

func TestRouterRegisterGroup(t *testing.T) {
	router := NewRouter()
	router.RegisterGroup(testRouteGroup{})

	request := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	response := httptest.NewRecorder()

	router.Handler().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if response.Body.String() != "test" {
		t.Fatalf("expected body %q, got %q", "test", response.Body.String())
	}
}
