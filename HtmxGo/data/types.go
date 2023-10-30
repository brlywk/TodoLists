package data

import "fmt"

// Struct representing a single todo item
type Todo struct {
	Id          int
	Name        string
	Description string
	Active      bool
	UserId		string
}

// Stringer method for Todo structs
func (todo Todo) String() string {
	return fmt.Sprintf("ID: %v,\tName: %v,\tDescription: %v,\tActive: %v\tUser: %v", todo.Id, todo.Name, todo.Description, todo.Active, todo.UserId)
}
