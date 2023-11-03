package data

import (
	"database/sql"
	"fmt"
	"log"
)

// Delete todo and return success or error
func DeleteTodoById(db *sql.DB, id int, userId string) (bool, error) {
	if userId == "" {
		return false, fmt.Errorf("UserId cannot be empty.")
	}

	_, err := db.Exec("DELETE FROM todos WHERE id = ? AND userId = ?",id, userId)
	if err != nil {
		log.Printf("\tDeleteTodoById\tDeletion Error\t%v", err)
		return false, err
	}

	return true, nil
}
