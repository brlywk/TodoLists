package data

import (
	"database/sql"
	"fmt"
	"log"
)

// Don't know if this works, but let's try it this way...
var (
	id          int
	name        string
	description string
	active      bool
	userId      string
)

// Queries a single Todo by ID
func GetSingleTodoById(db *sql.DB, id int) (Todo, error) {
	// if userId == "" {
	// 	return Todo{}, fmt.Errorf("UserId cannot be empty.")
	// }

	stmt, err := db.Prepare("SELECT * FROM todos WHERE id = ?")
	if err != nil {
		log.Printf("\tGetSingleTodoById\tPrepare Statement\t%s", err)
		return Todo{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&id, &name, &description, &active, &userId)
	if err != nil {
		log.Printf("\tGetSingleTodoById\tQueryRow\t%s", err)
		return Todo{}, err
	}

	newTodo := Todo{
		Id:          id,
		Name:        name,
		Description: description,
		Active:      active,
		UserId:      userId,
	}

	return newTodo, nil
}

// Return a list of all todos for a user, or an empty array & error
func GetAllTodosForUser(db *sql.DB, userId string) ([]Todo, error) {
	if userId == "" {
		return []Todo{}, fmt.Errorf("UserId cannot be empty.")
	}

	stmt, err := db.Prepare("SELECT * FROM todos WHERE userId = $1")
	if err != nil {
		log.Printf("\tGetAllTodosForUser\tPrepare Statement\t%s", err)
		return []Todo{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Printf("\tGetAllTodosForUser\tReceiving Rows\t%s", err)
		return []Todo{}, err
	}
	defer rows.Close()

	// return array
	allTodos := []Todo{}

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &active, &userId)
		if err != nil {
			log.Printf("\tGetAllTodosForUser\tSingle Row\t%s", err)
			continue
		}

		tmpTodo := Todo{
			Id:          id,
			Name:        name,
			Description: description,
			Active:      active,
			UserId:      userId,
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
