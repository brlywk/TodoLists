package api

import (
	"brlywk/HtmxGo/data"
	"brlywk/HtmxGo/utils"
	"net/http"
	"strconv"
)

// Deletes specified todo
// NOTE: For HTMX to delete an element an endpoint needs to return
// Status 200 with an empty body
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	rawTodoId := r.URL.Query().Get("id")
	userId := r.URL.Query().Get("userId")

	todoId, err := strconv.Atoi(rawTodoId)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	_, err = data.DeleteTodoById(data.DB, todoId, userId)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Return success but empty response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
