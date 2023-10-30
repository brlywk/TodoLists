package data

import (
	"database/sql"
	"log"
)

// Updates an existing todo. Returns if successful or error
func UpdateTodo(db *sql.DB, todo Todo) (bool, error) {
	stmt, err := db.Prepare("UPDATE todos SET name = $1, description = $2, active = $3 WHERE id = $4 AND userId = $5")
	if err != nil {
		log.Printf("\tUpdateTodo\tPrepare Statement\t%s", err)
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Name, todo.Description, todo.Active, todo.Id, todo.UserId)
	if err != nil {
		log.Printf("\tUpdateTodo\tUpdate Todo\t%s", err)
		return false, err
	}

	return true, nil
}
