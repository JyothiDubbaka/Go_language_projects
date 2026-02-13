package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func main() {
	//Check starting arguments
	if len(os.Args)!= 2 || os.Args[1] == "help" {
	printHelp()
	return
	}

	//Creating the file & preparing the terminal for the notestool
	collection := os.Args[1] + ".txt"
	notes := loadNotes(collection)
	reader := bufio.NewReader(os.Stdin)
	ClearScreen()
	fmt.Println(Bold + Green + "\nWelcome to the notes tool Folks! 🤝" + Reset)

	for {
		fmt.Println(Bold + Cyan + "\nSelect operation:" + Reset)
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
			//show contents
			case "1": 
				ClearScreen()
				showNotes(notes)

			//add new note
			case "2": 
				ClearScreen()
				fmt.Println("\nEnter the note text:")
				text,_ := reader.ReadString('\n')
				text = strings.TrimSpace(text)
					if text != "" {
						notes = append(notes, text)
						saveNotes(collection, notes)
					}

			//delete a note
			case "3":
				fmt.Println(Red + "\nEnter the number of note to remove or 0 to cancel:" + Reset)
				numStr,_ := reader.ReadString('\n')
				numStr = strings.TrimSpace(numStr)
				
				num, err := strconv.Atoi(numStr)
				if err != nil || num < 0 || num > len(notes) {
					fmt.Println("Non-existing note ID.")
					continue
				}
				if num == 0 {
					continue
				}
				notes = append(notes[:num-1], notes[num:]...)
				saveNotes(collection, notes)

			//Exiting the application
			case "4":
				fmt.Println(Bold + Green + "\nGoodbye 👋\n" + Reset)
				return

			//Check for invalid input 
			default:
				fmt.Println(Red + "\nInvalid option. Choose from 1-4." + Reset)
		}
	}

}
