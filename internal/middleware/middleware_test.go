package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestChainExecutionOrder(t *testing.T) {
	var execution []string

	first := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			execution = append(execution, "first-before")
			next.ServeHTTP(w, r)
			execution = append(execution, "first-after")
		})
	}

	second := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			execution = append(execution, "second-before")
			next.ServeHTTP(w, r)
			execution = append(execution, "second-after")
		})
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		execution = append(execution, "handler")
		w.WriteHeader(http.StatusOK)
	})

	chained := Chain(handler, first, second)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	chained.ServeHTTP(response, request)

	expected := []string{
		"first-before",
		"second-before",
		"handler",
		"second-after",
		"first-after",
	}

	if !reflect.DeepEqual(execution, expected) {
		t.Fatalf("expected execution order %v, got %v", expected, execution)
	}

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}
}
