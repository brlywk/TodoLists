package server

import (
	"brlywk/HtmxGo/utils"
	"fmt"
	"html/template"
	"net/http"
)

// This function really handles all routes possible
func GetRoot(w http.ResponseWriter, r *http.Request) {
	// NOTE: Super important to actually CALL the function 'measure' returns...
	defer utils.Measure(r.URL.Path, r.Method)()

	path := r.URL.Path

	pathMapping := map[string]string{
		"/":     "index.html",
		"/test": "test.html",
	}

	templateFile, exists := pathMapping[path]

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Page not found."))
		return
	}

	templatePath := fmt.Sprintf("templates/%s", templateFile)

	html, err := template.ParseFiles(templatePath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("File not found: %s", templateFile)))
		return
	}

	Test := struct {
		Name string
	}{
		Name: "Hello there",
	}

	html.Execute(w, Test)
}
