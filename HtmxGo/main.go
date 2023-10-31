package main

import (
	"brlywk/HtmxGo/api"
	"brlywk/HtmxGo/data"
	"brlywk/HtmxGo/server"
	"brlywk/HtmxGo/utils"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Nobody really knows what this function is supposed to do...
func main() {
	port := fmt.Sprintf(":%s", utils.GetEnv("PORT", "3000"))

	// prepare sqlite connection
	var err error
	data.DB, err = sql.Open("sqlite3", "./db/todo.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer data.DB.Close()

	// new server
	mux := http.NewServeMux()

	// server static files
	fs := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))


	// api handlers
	// NOTE: Go requires both of these routes in order to handle the route itself and
	// all subroutes of it...
	mux.HandleFunc("/api/todos", api.TodosRouter)
	mux.HandleFunc("/api/todos/", api.TodosRouter)

	// serve root files which basically is only the index.html
	mux.HandleFunc("/", server.GetRoot)

	// run server
	err = http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
