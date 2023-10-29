package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

// Helper function that tries to fetch an  env var, and if it fails, returns fallback
func getEnv(key string, fallback string) string {
	value, found := os.LookupEnv(key)

	if !found {
		return fallback
	}
	return value
}

// ---- simple template rendering -------------------------
// This function really handles all routes possible
func getRoot(w http.ResponseWriter, r *http.Request) {
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

// ---- setup for api request handling --------------------
type TestResponse struct {
	Name   string `json:"something_completely_different"`
	Answer int
}

func getTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Data := TestResponse{
		Name:   "Test JSON Response",
		Answer: 42,
	}
	json.NewEncoder(w).Encode(Data)
}

// Test handler to get partial HTML template
func getHtmlTest(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
		return
	}

	// To play around with HTMX loading a bit...
	time.Sleep(2 * time.Second)

	html, err := template.ParseFiles("templates/partials/htmxTest.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Template not found! OH NOES!"))
		return
	}

	html.Execute(w, "This has been added by the API function")
}

func main() {
	port := fmt.Sprintf(":%s", getEnv("PORT", "3000"))

	// new server
	mux := http.NewServeMux()

	// server static files
	fs := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	// serve root files which basically is only the index.html
	mux.HandleFunc("/", getRoot)

	// api handlers
	mux.HandleFunc("/api/test", getTest)
	mux.HandleFunc("/api/htmx", getHtmlTest)

	// run server
	err := http.ListenAndServe(port, mux)

	if err == nil {
		log.Fatal(err)
	}
}
