package main

import (
	"brlywk/HtmxGo/api"
	"brlywk/HtmxGo/data"
	"brlywk/HtmxGo/server"
	"brlywk/HtmxGo/utils"
	"os"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "github.com/mattn/go-sqlite3"
)

// Nobody really knows what this function is supposed to do...
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env file not found, trying to read from environment...")
	}

	port := fmt.Sprintf(":%s", utils.GetEnv("PORT", "8080"))
	dbUrl := os.Getenv("DB_URL")
	dbToken := os.Getenv("DB_TOKEN")
	fullDbUrl := fmt.Sprintf("%v?authToken=%v", dbUrl, dbToken)

	// prepare sqlite connection
	// data.DB, err = sql.Open("sqlite3", "./db/todo.sqlite3")
	data.DB, err = sql.Open("libsql", fullDbUrl)
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
	// GET
	mux.HandleFunc("/api/todos", api.GetApiTodos)
	mux.HandleFunc("/api/todo", api.GetTodoById)
	// POST
	mux.HandleFunc("/api/create", api.PostTodo)
	// PUT
	mux.HandleFunc("/api/changeUserId", api.GetChangeUserId)
	mux.HandleFunc("/api/toggle", api.PutToggleTodo)
	// DELETE
	mux.HandleFunc("/api/delete", api.DeleteTodo)

	// serve root files which basically is only the index.html
	mux.HandleFunc("/", server.GetRoot)

	// run server
	err = http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
