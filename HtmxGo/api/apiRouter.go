package api

import (
	"log"
	"net/http"
)

// Router handling all /todos routes
func TodosRouter(w http.ResponseWriter, r *http.Request) {
	log.Print("\t========== TodosRouter ==========")
	switch r.Method {
	// GET
	case http.MethodGet:
		GetApiTodos(w, r)
	// POST
	case http.MethodPost:
		// handler here
	case http.MethodDelete:
		// handler here
	case http.MethodPut:
		// handler here
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Unsupported Request Method"))
	}
}
