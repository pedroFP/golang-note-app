package main

import (
	"fmt"

	note "notesapp.com/app/model"
)

func main() {
	note.Initialize()

	// fmt.Println(note.Notes())

	newNote, err := note.New("New NOte title", false)
	if err != nil {
		fmt.Println("note not instanced")
		return
	}

	err = note.Save(newNote)
	if err != nil {
		fmt.Println("note not saved")
		return
	}
}
