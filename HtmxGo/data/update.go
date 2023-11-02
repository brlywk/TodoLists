package data

import (
	"database/sql"
	"log"
	"strconv"
)

// Updates an existing todo. Returns if successful or error
func UpdateTodo(db *sql.DB, todo Todo) (Todo, error) {
	fail := func(err error, message string) (Todo, error) {
		log.Printf("\tUpdateToggleTodo\t%v\t%v", message, err)
		return Todo{}, err
	}

	tx, err := db.Begin()
	if err != nil {
		return fail(err, "Unable to start transaction")
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE todos SET name = ? WHERE id = ?", todo.Name, id)
	if err != nil {
		return fail(err, "Error toggling active state for Todo with id "+strconv.Itoa(id))
	}

	var (
		tId         int
		name        string
		description string
		active      bool
		userId      string
	)

	row := tx.QueryRow("SELECT * FROM todos WHERE id = ?", todo.Id)
	err = row.Scan(&tId, &name, &description, &active, &userId)
	if err != nil {
		return fail(err, "Error while searching Todo with id "+strconv.Itoa(id))
	}

	tmpTodo := Todo{
		Id:          tId,
		Name:        name,
		Description: description,
		Active:      active,
		UserId:      userId,
	}

	if err := tx.Commit(); err != nil {
		return fail(err, "Unable to commit transaction")
	}

	return tmpTodo, nil
}

// Toggles the active state of a Todo with given id and returns the updated
// todo or an error
func UpdateToggleTodo(db *sql.DB, id int) (Todo, error) {
	fail := func(err error, message string) (Todo, error) {
		log.Printf("\tUpdateToggleTodo\t%v\t%v", message, err)
		return Todo{}, err
	}

	tx, err := db.Begin()
	if err != nil {
		return fail(err, "Unable to start transaction")
	}
	defer tx.Rollback()

	var (
		tId         int
		name        string
		description string
		active      bool
		userId      string
	)

	row := tx.QueryRow("SELECT * FROM todos WHERE id = ?", id)
	err = row.Scan(&tId, &name, &description, &active, &userId)
	if err != nil {
		return fail(err, "Error while searching Todo with id "+strconv.Itoa(id))
	}

	tmpTodo := Todo{
		Id:          id,
		Name:        name,
		Description: description,
		Active:      !active,
		UserId:      userId,
	}

	_, err = tx.Exec("UPDATE todos SET active = ? WHERE id = ?", !active, id)
	if err != nil {
		return fail(err, "Error toggling active state for Todo with id "+strconv.Itoa(id))
	}

	if err := tx.Commit(); err != nil {
		return fail(err, "Unable to commit transaction")
	}

	return tmpTodo, nil
}

// Update all entries for an old user id with a new id
// ... what do you mean with 'potential issues'?!
func UpdateUserId(db *sql.DB, oldUserId string, newUserId string) ([]Todo, error) {
	log.Printf("\tStarting DB stuff\nOld: %v\nNew: %v", oldUserId, newUserId)

	// Quick way to exit out on errors
	fail := func(err error, message string) ([]Todo, error) {
		log.Printf("\tUpdateUserId\t%v\t%v", message, err)
		return []Todo{}, err
	}

	// create transaction
	tx, err := db.Begin()
	if err != nil {
		return fail(err, "Unable to open transaction")
	}
	defer tx.Rollback()

	// Update rows
	_, err = tx.Exec("UPDATE todos SET userId = ? WHERE userId = ?", newUserId, oldUserId)
	if err != nil {
		return fail(err, "Error updating todos")
	}

	// let's try this for now, before requerying
	rows, err := tx.Query("SELECT * FROM todos WHERE userId = ?", newUserId)
	if err != nil {
		return fail(err, "Error getting rows")
	}
	defer rows.Close()

	// gotta slice 'em todos
	todos := []Todo{}

	for rows.Next() {
		// temp variables we need
		var (
			id          int
			name        string
			description string
			active      bool
			userId      string
		)

		err := rows.Scan(&id, &name, &description, &active, &userId)
		if err != nil {
			return fail(err, "Something went wrong converting row to Todo")
		}

		tmpTodo := Todo{
			Id:          id,
			Name:        name,
			Description: description,
			Active:      active,
			UserId:      userId,
		}

		todos = append(todos, tmpTodo)
	}
	err = rows.Err()
	if err != nil {
		return fail(err, "Error finishing iterating over rows")
	}

	// commit transaction if successful
	if err := tx.Commit(); err != nil {
		return fail(err, "Error committing transaction. It's rollback time!")
	}

	// delete
	return todos, err
}
