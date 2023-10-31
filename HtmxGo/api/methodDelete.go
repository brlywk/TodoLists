package api

import (
	"brlywk/HtmxGo/utils"
	"log"
	"net/http"
)

// Deletes specified todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)

	log.Printf("Delete Todo called with %v", r.URL.RawQuery)
}
