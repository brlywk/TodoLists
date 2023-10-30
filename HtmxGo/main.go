package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

// Helper function that tries to fetch an  env var, and if it fails, returns fallback
func getEnv(key string, fallback string) string {
	value, found := os.LookupEnv(key)

	if !found {
		return fallback
	}
	return value
}

// Really need getting used to defer and all of that...
func measure(name string, method string) func() {
	start := time.Now()

	if method == "" {
		method = "GET"
	}

	return func() {
		log.Printf("\tPath: %s\t\tMethod: %s\tTime: %v", name, method, time.Since(start))
	}
}

// ---- simple template rendering -------------------------
// This function really handles all routes possible
func getRoot(w http.ResponseWriter, r *http.Request) {
	// NOTE: Super important to actually CALL the function 'measure' returns...
	defer measure(r.URL.Path, r.Method)()

	path := r.URL.Path

	pathMapping := map[string]string{
		"/":     "index.html",
		"/test": "test.html",
	}

	templateFile, exists := pathMapping[path]

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Page not found."))
		return
	}

	templatePath := fmt.Sprintf("templates/%s", templateFile)

	html, err := template.ParseFiles(templatePath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("File not found: %s", templateFile)))
		return
	}

	Test := struct {
		Name string
	}{
		Name: "Hello there",
	}

	html.Execute(w, Test)
}

// ---- setup for api request handling --------------------
type TestResponse struct {
	Name   string `json:"something_completely_different"`
	Answer int
}

func getTest(w http.ResponseWriter, r *http.Request) {
	defer measure(r.URL.Path, r.Method)()

	w.Header().Set("Content-Type", "application/json")
	Data := TestResponse{
		Name:   "Test JSON Response",
		Answer: 42,
	}
	json.NewEncoder(w).Encode(Data)
}

// Test handler to get partial HTML template
func getHtmlTest(w http.ResponseWriter, r *http.Request) {
	// Just out of curiosity and for logging and such
	defer measure(r.URL.Path, r.Method)()

	if r.Method != http.MethodGet {
		return
	}

	// To play around with HTMX loading a bit...
	time.Sleep(2 * time.Second)

	html, err := template.ParseFiles("templates/partials/htmxTest.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Template not found! OH NOES!"))
		return
	}

	html.Execute(w, "This has been added by the API function")
}

// For SQL rows

type TestRow struct {
	id          int
	name        string
	description string
	active      bool
}

func main() {
	port := fmt.Sprintf(":%s", getEnv("PORT", "3000"))

	// prepare sqlite connection
	db, err := sql.Open("sqlite3", "./db/todo.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id          int
		name        string
		description string
		active      bool
	)

	// apparently it's better to prepare statements in order to reduce the total number of
	// calls made against the db
	// HOWEVER, this is only really true if a statement is reused multiple times 
	// (e.g. regular re-accessing of data)
	statement, err := db.Prepare("SELECT * FROM todos WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()


	// shorthand for querying a single row
	err = statement.QueryRow(2).Scan(&id, &name, &description, &active)
	if err != nil {
		// Best way to handle errors for queries is check whether any rows have been returned
		// at all
		if err == sql.ErrNoRows {
			log.Println("No rows returned :(")
		} else {
			log.Fatal(err)
		}
	}
	log.Printf("Result: %d\t%s\t%s\t%v", id, name, description, active)

	// // let's add a new test row
	// stmt, err := db.Prepare("INSERT INTO todos(name, description, active) VALUES ($1, $2, $3)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // we can use _ if we don't care for the statement result
	// _, err = stmt.Exec("Insert Test", "Let's thest creating a row via Go", 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()


	// ALL OF THE BELOW IS HOW TO QUERY MULTIPLE ROWS
	// run a test query
	// rows, err := db.Query("select * from todos where id = ?", 1)
	// if err != nil {
	// 	// just for testing, failed requests should not kill the server
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	//
	// for rows.Next() {
	// 	var (
	// 		id          int
	// 		name        string
	// 		description string
	// 		active      bool
	// 	)
	//
	// 	err := rows.Scan(&id, &name, &description, &active)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	// TODO: Here we would create a new TestRow and move the values in probably?
	// 	log.Printf("Result: %d\t%s\t%s\t%v", id, name, description, active)
	// }
	// // we need to check if something happend in the loopdiloop
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// new server
	mux := http.NewServeMux()

	// server static files
	fs := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	// serve root files which basically is only the index.html
	mux.HandleFunc("/", getRoot)

	// api handlers
	mux.HandleFunc("/api/test", getTest)
	mux.HandleFunc("/api/htmx", getHtmlTest)

	// run server
	err = http.ListenAndServe(port, mux)

	if err == nil {
		log.Fatal(err)
	}
}
