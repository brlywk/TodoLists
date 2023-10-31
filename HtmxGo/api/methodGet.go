package api

import (
	"log"
	"strconv"
	"time"

	"html/template"
	"net/http"

	"brlywk/HtmxGo/data"
	"brlywk/HtmxGo/utils"
)

// Route handler
//
//	GET
//		/api/todos/:user
//		/api/todos/:user/:id
func GetApiTodos(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	// TODO: Delay added for testing
	time.Sleep(2 * time.Second)

	// Extract necessary info from query string
	userId := r.URL.Query().Get("userId")
	todoIdStr := r.URL.Query().Get("todoId")

	// Restrict templates to api subfolder
	apiFS, err := getApiFS()
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Unable to access templates. Please try again.")
		return
	}

	// load templates
	// ... and as everyone knows, hardcoding absolutely rules and has no negative
	// side-effects whatsoever! /s
	templ, err := template.ParseFS(apiFS, "todo.html", "todolist.html")
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to parse template files")
		return
	}

	// 1 param  -> get all for user
	if userId != "" && todoIdStr == "" {
		userTodos, err := data.GetAllTodosForUser(data.DB, userId)
		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, "Unable to get todos for user")
			return
		}

		templ.ExecuteTemplate(w, "todoList", userTodos)
		return
	}

	// 2 params -> get id for user
	if userId != "" && todoIdStr != "" {
		todoId, _ := strconv.Atoi(todoIdStr)

		userTodo, err := data.GetSingleTodoByUserId(data.DB, todoId, userId)
		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, "Unable to get todo for user")
			return
		}

		templ.ExecuteTemplate(w, "todoItem", userTodo)
		return
	}
}

// Returns edit form for the specified todo item
func GetTodoEditForm(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	log.Printf("GetTodoEditForm called with %v", r.URL.RawQuery)
}

// // Just a simple test handler that returns a JSON object
// func GetTest(w http.ResponseWriter, r *http.Request) {
// 	defer utils.Measure(r.URL.Path, r.Method)()
//
// 	w.Header().Set("Content-Type", "application/json")
// 	Data := TestResponse{
// 		Name:   "Test JSON Response",
// 		Answer: 42,
// 	}
// 	json.NewEncoder(w).Encode(Data)
// }
