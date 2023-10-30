package data

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteTodoById(db *sql.DB, id int, userId string) (bool, error) {
	if userId == "" {
		return false, fmt.Errorf("UserId cannot be empty.")
	}

	stmt, err := db.Prepare("DELETE FROM todos WHERE id = $1 AND userId = $2")
	if err != nil {
		log.Printf("\tDeleteTodoById\tPrepare Statement\t%s", err)
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, userId)
	if err != nil {
		log.Printf("\tDeleteTodoById\tDeletion Error\t%v", err)
		return false, err
	}

	return true, nil
}
