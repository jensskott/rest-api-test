package main

import "fmt"

var currentID int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Install docker frameworks"})
	RepoCreateTodo(Todo{Name: "Choose between rancher/kubernetes"})
}

// RepoFindTodo to be exported to handlers.go
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

//RepoCreateTodo This is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.Id = currentID
	todos = append(todos, t)
	return t
}

//RepoDestroyTodo Hope this works, havnt tried it
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
