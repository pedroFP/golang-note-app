package note

import (
	"encoding/json"
	"fmt"
	"os"
)

const notesFileName = "notes.json"

type note struct {
	id         int
	title      string
	inProgress bool
}

func Initialize() error {
	file, err := os.OpenFile(notesFileName, os.O_CREATE|os.O_RDWR, 0o644)
	fmt.Println("File initialized", file)

	notes, err := Notes()

	if err != nil || len(notes) == 0 {
		err := os.WriteFile(notesFileName, []byte("[empty file]"), 0o644)
		return err
	}

	return nil
}

func New(title string, inProgress bool) (note, error) {
	id := generateID()
	return note{
		id:         id,
		title:      title,
		inProgress: inProgress,
	}, nil
}

func Save(n note) error {
	json, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error")
		return err
	}

	return os.WriteFile(notesFileName, json, 0o644)
}

func Notes() ([]note, error) {
	rawData, err := os.ReadFile(notesFileName)
	if err != nil {
		fmt.Println("could not read from file: ", notesFileName)
		return nil, err
	}

	var notes []note

	err = json.Unmarshal(rawData, &notes)
	if err != nil {
		fmt.Println("Error Unmarshalling JSON:", err)
		return nil, err
	}

	return notes, nil
}

func generateID() int {
	// TODO: Should be the count of total notes + 1 | consider a notes index in cache (a variable)
	// TODO: Count the total notes based on the length of the list read from the JSON file
	return 0
}
