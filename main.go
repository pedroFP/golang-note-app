package main

import (
	"fmt"

	note "notesapp.com/app/model"
)

func main() {
	err := note.Initialize()
	if err != nil {
		fmt.Println("note not instanced")
		return
	}

	newNote, err := note.New("New NOte title 2", false)
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
	fmt.Println(notes)
}
