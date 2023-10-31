package api

import (
	"log"
	"strconv"
	"strings"
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
	log.Print("\t---- GET Request ----")

	// NOTE:
	//  We have to handle the following routes
	//  /api/todos               Nothing for GET
	//  /api/todos/:userId       Get all Todos for userId
	//  /api/todos/:userId/:id   Get todo with id for userId

	path := r.URL.Path

	_, action, found := strings.Cut(path, "todos")
	if !found {
		WriteErrorResponse(w, http.StatusInternalServerError, "Unexpected error tokenizing request URL.")
		return
	}

	// remove leading and trailing / for further processing
	action, _ = strings.CutPrefix(action, "/")
	action, _ = strings.CutSuffix(action, "/")

	splitAction := strings.Split(action, "/")
	userId := splitAction[0]
	todoId := -1

	// we don't care about any index > 1
	if len(splitAction) > 1 {
		var err error
		todoId, err = strconv.Atoi(splitAction[1])
		if err != nil {
			// supremely disappointing that Go overwrites an int with the 'default'
			// it the conversion fails... makes absolute sense, but super inconvenient right now!
			todoId = -1
		}
	}

	// nothing to do on 'top level' requests
	if action == "" {
		WriteErrorResponse(w, http.StatusNotFound, "This endpoint provides no functionality.")
		return
	}

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
	if userId != "" && todoId == -1 {
		userTodos, err := data.GetAllTodosForUser(data.DB, userId)
		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, "Unable to get todos for user")
			return
		}

		templ.ExecuteTemplate(w, "todoList", userTodos)
		return
	}

	// 2 params -> get id for user
	if userId != "" && todoId > -1 {
		userTodo, err := data.GetSingleTodoByUserId(data.DB, todoId, userId)
		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, "Unable to get todo for user")
			return
		}

		templ.ExecuteTemplate(w, "todoItem", userTodo)
		return
	}
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

	html, err := template.ParseFiles("templates/api/htmxTest.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Template not found! OH NOES!"))
		return
	}

	html.Execute(w, "This has been added by the API function")
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
