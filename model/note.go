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
	return note{
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

	n.ID = len(notes) + 1

	updatedNotes := append(notes, *n)

	data, err := json.Marshal(updatedNotes)
	if err != nil {
		fmt.Println("Error on Marshal updatedNotes")
		return err
	}

	return os.WriteFile(notesFileName, data, 0o644)
}

func (n *note) Update() error {
	notes, err := Notes()
	if err != nil {
		return err
	}

	for index, element := range notes {
		if element.ID == n.ID {
			notes[index] = *n
			break
		}
	}

	data, err := json.Marshal(notes)
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
