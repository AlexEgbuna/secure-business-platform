package api

import (
	"encoding/json"
	"net/http"
)

// Version identifies the current application build version.
//
// The default value is intentionally simple for the development stage.
// A future build/release process can inject the production version.
var Version = "dev"

type versionResponse struct {
	Version string `json:"version"`
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := versionResponse{
		Version: Version,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
