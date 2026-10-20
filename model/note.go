package note

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const notesFileName = "notes.json"

type note struct {
	ID         int
	Title      string
	InProgress bool
}

func Initialize() error {
	file, err := os.OpenFile(notesFileName, os.O_CREATE|os.O_RDWR, 0o644)
	fmt.Println("File initialized open or created")

	if err != nil {
		return err
	}

	defer file.Close()

	var notes []note

	err = json.NewDecoder(file).Decode(&notes)

	if (err != nil && err != io.EOF) || len(notes) == 0 {
		file.Truncate(0)
		file.Seek(0, 0)
		_, err = file.Write([]byte("[]"))
		fmt.Println("File was set to '[]'")
		return err
	}

	fmt.Println("File already exists")
	return nil
}

func New(title string, inProgress bool) (note, error) {
	id := generateID()
	return note{
		ID:         id,
		Title:      title,
		InProgress: inProgress,
	}, nil
}

func (n *note) Save() error {
	notes, err := Notes()
	if err != nil {
		fmt.Println("Error on save")
		return err
	}

	updatedNotes := append(notes, *n)

	data, err := json.Marshal(updatedNotes)
	if err != nil {
		fmt.Println("Error on Marshal updatedNotes")
		return err
	}

	return os.WriteFile(notesFileName, data, 0o644)
}

func Notes() ([]note, error) {
	file, err := os.OpenFile(notesFileName, os.O_RDONLY, 0o644)
	if err != nil {
		return nil, err
	}

	var notes []note

	err = json.NewDecoder(file).Decode(&notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func generateID() int {
	// TODO: Should be the count of total notes + 1 | consider a notes index in cache (a variable)
	// TODO: Count the total notes based on the length of the list read from the JSON file
	return 0
}
