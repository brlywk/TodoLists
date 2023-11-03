package data

import (
	"database/sql"
	"fmt"
	"log"
)

// Queries a single Todo by ID
func GetSingleTodoById(db *sql.DB, id int) (Todo, error) {
	var (
		tId     int
		tName   string
		tDesc   string
		tActive bool
		tUserId string
	)

	err := db.QueryRow("SELECT * FROM todos WHERE id = ?", id).Scan(&tId, &tName, &tDesc, &tActive, &tUserId)
	if err != nil {
		log.Printf("\tGetSingleTodoById\tQueryRow\t%s", err)
		return Todo{}, err
	}

	tmpTodo := Todo{
		Id:          tId,
		Name:        tName,
		Description: tDesc,
		Active:      tActive,
		UserId:      tUserId,
	}

	return tmpTodo, nil
}

// Return a list of all todos for a user, or an empty array & error
func GetAllTodosForUser(db *sql.DB, userId string) ([]Todo, error) {

	var (
		tId     int
		tName   string
		tDesc   string
		tActive bool
		tUserId string
	)

	if userId == "" {
		return []Todo{}, fmt.Errorf("UserId cannot be empty.")
	}

	rows, err := db.Query("SELECT * FROM todos WHERE userId = ?", userId)
	if err != nil {
		log.Printf("\tGetAllTodosForUser\tReceiving Rows\t%s", err)
		return []Todo{}, err
	}
	defer rows.Close()

	// return array
	allTodos := []Todo{}

	for rows.Next() {
		err := rows.Scan(&tId, &tName, &tDesc, &tActive, &tUserId)
		if err != nil {
			log.Printf("\tGetAllTodosForUser\tSingle Row\t%s", err)
			continue
		}

		tmpTodo := Todo{
			Id:          tId,
			Name:        tName,
			Description: tDesc,
			Active:      tActive,
			UserId:      tUserId,
		}

		allTodos = append(allTodos, tmpTodo)
	}
	// check if any errors in loop happened
	if err = rows.Err(); err != nil {
		log.Printf("\tGetAllTodosForUser\tIterating Rows\t%s", err)
		return []Todo{}, err
	}

	return allTodos, nil
}
