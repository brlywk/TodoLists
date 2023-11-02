package api

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
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

// Get the requested API templates
func GetApiTemplates(files ...string) (*template.Template, error) {
	// get templates from FS
	apiFS, err := getApiFS()
	if err != nil {
		return nil, err
	}

	templ, err := template.ParseFS(apiFS, files...)
	if err != nil {
		return nil, err
	}

	return templ, nil
}

// Simple helper to save some writing on sending an error response
// Currently only accepts errors and strings as message
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	var msg string
	switch message.(type) {
	case error:
		msg = message.(error).Error()
	case string:
		msg = message.(string)
	default:
		msg = fmt.Sprintf("Unsupported type %t", message)
	}

	log.Printf("\tStatus: %v\tMessage: %v", statusCode, msg)
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
	return
}
