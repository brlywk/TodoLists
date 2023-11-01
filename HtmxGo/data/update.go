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

		log.Printf("Todo added: %v", tmpTodo)

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

	log.Print("And we're done, exiting transaction")

	// delete
	return todos, err
}
