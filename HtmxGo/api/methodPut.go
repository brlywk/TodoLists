package api

import (
	"brlywk/HtmxGo/utils"
	"log"
	"net/http"
)

// Toggles active state of a todo
func PutToggleTodo(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)

	log.Printf("PutToggleData called with Query %v", r.URL.RawQuery)
}

// Finished an Edit and saves changes
func PutEditTodo(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)

	log.Printf("PutEditTodo called with %v", r.URL.RawQuery)
}
