package main

import (
	"log"

	sdk "github.com/TyunTech/go-todostore"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	config := &sdk.Config{
		// Address: "insert-your-todostore-address-here",
		Address: "http://localhost:8000",
	}

	client, err := sdk.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// List all Todos
	todoL, _ := client.Todos.List(sdk.TodoListOptions{})
	spew.Printf("【todoList】: %v\n", todoL)

	// Create a new todo
	options := sdk.TodoCreateOptions{
		Title: "Golang Todo",
		Memo:  "Write a Golang client for TodoStore",
	}

	todo, err := client.Todos.Create(options)
	if err != nil {
		log.Fatal(err)
	}

	// debug output
	spew.Printf("【debug todo id】: %v \n", todo.ID)
	// Update todo
	todo, _ = client.Todos.Update(todo.ID, sdk.TodoUpdateOptions{Title: "Golang Todo modified", Memo: "haha"})
	spew.Printf("【updated_todo】: %v \n", todo)

	// Read todo by ID
	todo, _ = client.Todos.Read(todo.ID)
	spew.Printf("【todo】: %v \n", todo)

	// Delete a todo
	err = client.Todos.Delete(todo.ID)
	if err != nil {
		log.Fatal(err)
	}

	// // List all Todos
	todoL, _ = client.Todos.List(sdk.TodoListOptions{})
	spew.Printf("todoList: %v\n", todoL)
}
