package main

import (
	"fmt"
	"bufio"
	"os"
)

//Printing help message
func printHelp() {
	fmt.Println("Usage: Run the application by ./notestool [Collection_Name]")
}

//Creating the file if it is not there and opening it if it is there.
func loadNotes(filename string) []string {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		os.Exit(1)
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}
	return notes
}

//Writing in the file and saving it.
func saveNotes(filename string, notes []string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error saving notes", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, note :=range notes {
		writer.WriteString(note + "\n")
	}
	writer.Flush()
}

//Showing the contents of the file
func showNotes(notes []string) {
	if len(notes) == 0 {
		fmt.Println(Bold + Yellow + "\nNotes:" + Reset + "\n(no notes)")
		return
	}
	fmt.Println(Bold + Yellow + "\nNotes:" + Reset)
	for i, note := range notes {
		fmt.Printf("%03d - %s\n",i+1,note)
	}
}