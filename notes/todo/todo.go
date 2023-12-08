package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (t Todo) Display() {
	fmt.Println(t.Text)
}

func (n Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		errorText := "invalid input"
		return Todo{}, errors.New(errorText)
	}

	return Todo{
		content,
	}, nil
}
