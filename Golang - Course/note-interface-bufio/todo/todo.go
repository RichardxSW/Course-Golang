package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	// "strings"
	// "time"
)

type Todo struct {
	// Title     string    `json:"title"`
	Text string `json:"text"`
	// CreatedAt time.Time `json:"created_at"`
}

func (t Todo) Display() {
	fmt.Println(t.Text)
}

// tidak pakai pointer ( * ) karena tidak ada perubahan data
func (t Todo) Save() error {
	// fileName := strings.ReplaceAll(n.Title, " ", "_")
	// fileName = strings.ToLower(fileName) + ".json"
	fileName := "todo.json"

	json, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
// return Todo ga pake * karena tidak ada perubahan
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Note must have a title and content")
	}

	return Todo{
		// Title:     title,
		Text:   content,
		// CreatedAt: time.Now(),
	}, nil
}
