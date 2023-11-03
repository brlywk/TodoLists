package api

import (
	"html/template"
	"strconv"

	"net/http"

	"brlywk/HtmxGo/data"
	"brlywk/HtmxGo/utils"
)

// Returns either all todos for a user, or a single todo,
// depending on wheter query parameter 'id' is provided
func GetApiTodos(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	// Extract necessary info from query string
	userId := r.URL.Query().Get("userId")
	todoIdStr := r.URL.Query().Get("todoId")

	// load templates
	// ... and as everyone knows, hardcoding absolutely rules and has no negative
	// side-effects whatsoever! /s
	templ, err := GetApiTemplates("todo.html", "todolist.html")
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// 1 param  -> get all for user
	if userId != "" && todoIdStr == "" {
		userTodos, err := data.GetAllTodosForUser(data.DB, userId)
		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		templ.ExecuteTemplate(w, "todoList", userTodos)
		return
	}

	// 2 params -> get id for user
	// NOTE: This one changed, so now it does not really need a user ID anymore
	if userId != "" && todoIdStr != "" {
		todoId, _ := strconv.Atoi(todoIdStr)

		userTodo, err := data.GetSingleTodoById(data.DB, todoId)
		if err != nil {
			WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		templ.ExecuteTemplate(w, "todoItem", userTodo)
		return
	}
}

// Used to update an existing user id with a new user id and refetch
// all todos for the updated id. this makes it easier to work
// with HTXM
func GetChangeUserId(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	oldUserId := r.URL.Query().Get("old")
	newUserId := r.URL.Query().Get("new")

	updatedTodos, err := data.UpdateUserId(data.DB, oldUserId, newUserId)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// load templates
	// ... and as everyone knows, hardcoding absolutely rules and has no negative
	// side-effects whatsoever! /s
	templ, err := GetApiTemplates("todo.html", "todolist.html")
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	templ.ExecuteTemplate(w, "todoList", updatedTodos)
}

// Return single todo item
// If action = present is in query send form, otherwise rendered todo
func GetTodoById(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()

	rawTodoId := r.URL.Query().Get("id")
	action := r.URL.Query().Get("action")

	todoId, err := strconv.Atoi(rawTodoId)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	var todo data.Todo

	// if action is 'save' we need to update and return the
	// updated todo
	if action == "save" {
		err := r.ParseForm()
		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		userId := r.PostFormValue("userId")
		name := r.PostFormValue("name")
		strActive := r.PostFormValue("active")

		if userId == "" || name == "" || strActive == "" {
			WriteErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		active, err := strconv.ParseBool(strActive)
		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, err)
		}

		tmpTodo := data.Todo{
			Id:          todoId,
			Name:        name,
			Description: "",
			Active:      active,
			UserId:      userId,
		}

		todo, err = data.UpdateTodo(data.DB, tmpTodo)
	} else {
		todo, err = data.GetSingleTodoById(data.DB, todoId)
	}
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// depending on action get edit or display template
	var tmpl *template.Template
	var tErr error
	tmplName := ""

	// for both show (no action) and 'save' we want
	// to return the display template
	if action == "edit" {
		tmpl, tErr = GetApiTemplates("editTodo.html")
		tmplName = "editTodoItem"
	} else {
		tmpl, tErr = GetApiTemplates("todo.html")
		tmplName = "todoItem"
	}
	if tErr != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.ExecuteTemplate(w, tmplName, todo)
}
