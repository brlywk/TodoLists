package server

import (
	"fmt"
	"log"
	"strings"

	"html/template"
	"net/http"

	"brlywk/HtmxGo/templates"
	"brlywk/HtmxGo/utils"
)

// This function really handles all (non API) routes possible
func GetRoot(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	path := r.URL.Path

	// We don't want to handle any /api routes here
	if strings.Contains(path, "/api") {
		log.Print("\tAPI route detected, skipping")
		return
	}

	// Parse all templates from our embeded FS
	templ, err := template.ParseFS(&templates.Files, "*.html", "partials/*.html", "api/newTodoForm.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse template files"))
		return
	}

	err = templ.ExecuteTemplate(w, "htmlBase", pageData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Template Rendering Error: %v", err)))
		return
	}
}
