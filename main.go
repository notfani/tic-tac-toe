package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tic-tac-toe/game"
)

func main() {
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Welcome to Tic-Tac-Toe!")
		fmt.Println("1. Start a new game")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")

		var firstInput string
		firstInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		firstInput = strings.TrimSpace(firstInput)

		if firstInput == "2" {
			break
		}

		for {
			if game.InitBoard() {
				break
			}
		}

		game.Play()
	}

}
