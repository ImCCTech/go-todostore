package todostore

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"
)

// Compile-time proof of interface implementation.
var _ Todos = (*todos)(nil)
var ctx = context.Background()

// Todos describes all the todo related methods that the Todostore
// API supports.
type Todos interface {
	// List all the todos.
	List(options TodoListOptions) (*TodoList, error)

	// Create a new todo with the given options.
	Create(options TodoCreateOptions) (*Todo, error)

	// Read an todo by its ID.
	Read(todoID string) (*Todo, error)

	// Update an todo by its ID.
	Update(todoID string, options TodoUpdateOptions) (*Todo, error)

	// Delete an todo by its ID.
	Delete(todoID string) error
}

// todos implements Todos.
type todos struct {
	client *Client
}

// TodoList represents a list of todos.
type TodoList struct {
	Items []*Todo `json:"items"`
}

// Todo represents a Todostore todo.
type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Memo      string    `json:"memo"`
	Created   time.Time `json:"created"`
	Completed bool      `json:"completed"`
}

// TodoListOptions represents the options for listing todos.
type TodoListOptions struct {
	Limit int
}

// List all todos.
func (s *todos) List(options TodoListOptions) (*TodoList, error) {
	if options.Limit == 0 {
		options.Limit = 100
	}
	req, err := s.client.newRequest("GET", "todos", &options)
	if err != nil {
		return nil, err
	}

	todol := &TodoList{}
	err = s.client.do(ctx, req, todol)
	if err != nil {
		return nil, err
	}

	return todol, nil
}

// TodoCreateOptions represents the options for creating an todo.
type TodoCreateOptions struct {
	Title string `json:"title"`
	Memo  string `json:"memo"`
}

func (p TodoCreateOptions) valid() error {
	if !validString(&p.Title) {
		return errors.New("todo title is required")
	}
	if !validString(&p.Memo) {
		return errors.New("todo memo is required")
	}
	return nil
}

// Create a new todo with the given options.
func (s *todos) Create(options TodoCreateOptions) (*Todo, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	req, err := s.client.newRequest("POST", "todos", &options)
	if err != nil {
		return nil, err
	}

	ent := &Todo{}
	err = s.client.do(ctx, req, ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

// Read an todo by id.
func (s *todos) Read(todoID string) (*Todo, error) {
	if !validStringID(&todoID) {
		return nil, errors.New("invalid value for todoID")
	}

	u := fmt.Sprintf("todos/%s", url.QueryEscape(todoID))
	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	ent := &Todo{}
	err = s.client.do(ctx, req, ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

// TodoUpdateOptions represents the options for updating an todo.
type TodoUpdateOptions struct {
	Title string `json:"title"`
	Memo  string `json:"memo"`
}

// Update attributes of an existing todo.
func (s *todos) Update(todoID string, options TodoUpdateOptions) (*Todo, error) {
	if !validStringID(&todoID) {
		return nil, errors.New("invalid value for todoID")
	}

	u := fmt.Sprintf("todos/%s", url.QueryEscape(todoID))
	req, err := s.client.newRequest("PATCH", u, &options)
	if err != nil {
		return nil, err
	}

	ent := &Todo{}
	err = s.client.do(ctx, req, ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

// Delete an tood by its ID.
func (s *todos) Delete(todoID string) error {
	if !validStringID(&todoID) {
		return errors.New("invalid value for todoID")
	}

	u := fmt.Sprintf("todos/%s", url.QueryEscape(todoID))
	req, err := s.client.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return s.client.do(ctx, req, nil)
}
