package api

import (
	"io/fs"
	"net/http"

	"brlywk/HtmxGo/templates"
)

// Helper: Returns a FS restricted to only 'api' files
func getApiFS() (fs.FS, error) {
	apiFS, err := fs.Sub(&templates.Files, "api")
	if err != nil {
		return nil, err
	}

	return apiFS, nil
}

// Simple helper to save some writing on sending an error response
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
	return
}