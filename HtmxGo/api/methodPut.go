package api

import (
	"brlywk/HtmxGo/data"
	"brlywk/HtmxGo/utils"
	"log"
	"net/http"
	"strconv"
)

// Toggles active state of a todo
func PutToggleTodo(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	rawTodoId := r.URL.Query().Get("id")

	todoId, err := strconv.Atoi(rawTodoId)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Error: ID must be a number")
		return
	}

	updatedTodo, err := data.UpdateToggleTodo(data.DB, todoId)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Error while trying to update todo. Better luck next time.")
		return
	}

	tmpl, err := GetApiTemplates("todo.html")
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Error fetching todo template.")
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.ExecuteTemplate(w, "todoItem", updatedTodo)
}

// Finished an Edit and saves changes
func PutEditTodo(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	log.Printf("PutEditTodo called with %v", r.URL.RawQuery)
}
