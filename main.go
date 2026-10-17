package main

import (
	"fmt"

	note "notesapp.com/app/model"
)

func main() {
	note.Initialize()

	// return

	notes, err := note.Notes()
	if err != nil {
		fmt.Println("Error on Initialize: ", err)
		return
	}

	fmt.Println(notes)

	newNote, err := note.New("New NOte title", false)
	if err != nil {
		fmt.Println("note not instanced")
		return
	}

	err = newNote.Save()
	if err != nil {
		fmt.Println("note not saved")
		return
	}

	notes, _ = note.Notes()
	fmt.Println(notes)
}
