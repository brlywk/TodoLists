package data

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

// Creates a new Todo. Returns whether the Todo was created (bool) / an error occured
func CreateNewTodo(db *sql.DB, newName string, forUser string) (bool, error) {
	stmt, err := db.Prepare("INSERT INTO todos(name, userId) VALUES ($1, $2)")
	if err != nil {
		log.Printf("\tCreateNewTodo\tPreprate Statement\t%v", err)
		return false, err
	}
	defer stmt.Close()

	// check if userId is given, otherwise this is for a new user
	if forUser == "" {
		forUser = uuid.NewString() 
	}

	_, err = stmt.Exec(newName, forUser)
	if err != nil {
		log.Printf("\tCreateNewTodo\tExecute Statement\t%v", err)
		return false, err
	}

	return true, nil
}
