package api

import (
	"brlywk/HtmxGo/data"
	"brlywk/HtmxGo/utils"
	"net/http"
)

// Create a new todo
func PostTodo(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)()
	
	// parse form data 
	err := r.ParseForm()
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Unable to parse form data posted")
		return
	}

	userId := r.PostFormValue("userId")
	name := r.PostFormValue("name")	

	if userId == "" || name == "" {
		WriteErrorResponse(w, http.StatusInternalServerError, "Error: userId and name, must both be provided and must not be empty")
		return
	}

	_, err = data.CreateNewTodo(data.DB, name, userId)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Error creating new todo")
		return
	}

	tmpl, err := GetApiTemplates("newTodoForm.html")
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Something went wrong with the form. That should not have happened. We are doomed... DOOMED!")
		return
	}

	w.Header()["HX-Trigger"] = []string{"triggerLoad"}
	w.WriteHeader(http.StatusCreated)
	err = tmpl.ExecuteTemplate(w, "newTodoForm", nil)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Error rendering template")
		return
	}
}
