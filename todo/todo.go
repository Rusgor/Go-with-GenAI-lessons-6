// Package todo provides functions to persist a simple todo list to disk
// as JSON.
package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Todo represents a single todo-list item.
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// SaveTodos writes the given todos to path as indented JSON, creating the
// file if needed and replacing its contents if it already exists.
func SaveTodos(path string, todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal todos: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write todos to %q: %w", path, err)
	}
	return nil
}

// LoadTodos reads and parses the todo list stored at path. Wrapped errors
// retain their underlying causes for errors.Is and errors.As.
func LoadTodos(path string) ([]Todo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read todos from %q: %w", path, err)
	}

	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("decode todos from %q: %w", path, err)
	}
	return todos, nil
}
