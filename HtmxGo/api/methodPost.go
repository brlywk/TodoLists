package api

import (
	"brlywk/HtmxGo/utils"
	"log"
	"net/http"
)

func PostTodo(w http.ResponseWriter, r *http.Request) {
	defer utils.Measure(r.URL.Path, r.Method)

	log.Print("PostTodo called")
}
