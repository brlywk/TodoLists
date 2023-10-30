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
	db, err := sql.Open("sqlite3", "./db/todo.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// todo, err := data.GetSingleTodoByUserId(db, 1, "void")
	// if err != nil {
	// 	// Best way to handle errors for queries is check whether any rows have been returned
	// 	// at all
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No rows returned :(")
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }
	// log.Printf("\tTodo found:\t%v", todo)

	// Test insertion
	// success, err := data.CreateNewTodo(db, "Test Insertion", "", true)
	// if success {
	// 	log.Printf("Successfully created new Todo")
	// }

	// Get all rows
	// todos, err := data.GetAllTodosForUser(db, "void")
	// if err != nil {
	// 	log.Printf("Error: %v", err)
	// } else {
	// 	log.Printf("Todos: %v", todos)
	// }

	// success, err := data.DeleteTodoById(db, 4, "void")
	// if success {
	// 	log.Printf("Successfully deleted")
	// }

	// tmpTodo := data.Todo{
	// 	Id:          3,
	// 	Name:        "New Name",
	// 	Description: "New Description",
	// 	Active:      false,
	// 	UserId:      "void",
	// }
	// success, err := data.UpdateTodo(db, tmpTodo)
	// if success {
	// 	log.Printf("Todo with id %v updated", tmpTodo.Id)
	// }

	// new server
	mux := http.NewServeMux()

	// server static files
	fs := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	// serve root files which basically is only the index.html
	mux.HandleFunc("/", server.GetRoot)

	// api handlers
	mux.HandleFunc("/api/test", api.GetTest)
	mux.HandleFunc("/api/htmx", api.GetHtmlTest)

	// run server
	err = http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
