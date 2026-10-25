package main

import (
	"encoding/json"
	"fmt"

	note "notesapp.com/app/model"
)

func main() {
	err := note.Initialize()
	if err != nil {
		fmt.Println("note not instanced")
		return
	}

	newNote, err := note.New("HELLOW WORLD", false)
	if err != nil {
		fmt.Println("note not instanced")
		return
	}

	err = newNote.Save()
	if err != nil {
		fmt.Println("note not saved")
		return
	}

	notes, _ := note.Notes()

	buffer, _ := json.Marshal(notes)
	fmt.Println(string(buffer))

	// Upate newNote name

	newNote.Title = "Resident Evil"

	newNote.Update()

	notes, _ = note.Notes()

	buffer, _ = json.Marshal(notes)
	fmt.Println(string(buffer))
}
