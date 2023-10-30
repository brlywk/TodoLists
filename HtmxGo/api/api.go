package api

import (
	"brlywk/HtmxGo/utils"
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

// ---- setup for api request handling --------------------
type TestResponse struct {
	Name   string `json:"something_completely_different"`
	Answer int
}

// Just a simple test handler that returns a JSON object
func GetTest(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	w.Header().Set("Content-Type", "application/json")
	Data := TestResponse{
		Name:   "Test JSON Response",
		Answer: 42,
	}
	json.NewEncoder(w).Encode(Data)
}

// Test handler to get partial HTML template
func GetHtmlTest(w http.ResponseWriter, r *http.Request) {
	// Just out of curiosity and for logging and such
	defer utils.Measure(r.URL.Path, r.Method)()

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
