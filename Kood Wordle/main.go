package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"koodWordle/game"
	"koodWordle/io"
	"koodWordle/model"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <word-index>")
		return
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil || index < 0 {
		fmt.Println("Invalid word index")
		return
	}

	words, err := io.LoadWords("wordle-words.txt")
	if err != nil || index >= len(words) {
		fmt.Println("Word list missing or index out of range")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your username:")
	if !scanner.Scan() {
		return
	}
	user := model.NewUser(scanner.Text())

	game.Play(scanner, words[index], user)
}
