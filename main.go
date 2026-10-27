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

	newNote, err := note.New("My New Setup", false)
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

	printJSON(notes)

	note, err := notes.Find(2)
	if err != nil {
		fmt.Println(err)
	}

	printJSON(note)
}

func printJSON(note any) {
	buffer, _ := json.Marshal(note)
	fmt.Println("Note found!", string(buffer))
}
