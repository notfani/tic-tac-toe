package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GamesState int
type BoardField int

const (
	empty BoardField = iota
	cross
	nought
)

const (
	playing GamesState = iota
	draw
	crossWin
	noughtWin
	quit
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

		boardSize := 3 // Размер доски 3x3 (по-умолчанию)
		state := playing
		currentPlayer := cross

		for {
			fmt.Print("Enter the size of the board (3-9): ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			input = strings.TrimSpace(input)
			size, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Error reading input. Please try again.")
				continue
			}

			if size < 3 || size > 9 {
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			boardSize = size
			break
		}

		// Инициализация доски
		board := make([][]BoardField, boardSize)
		for i := range boardSize {
			board[i] = make([]BoardField, boardSize)
		}

		// Вывод в терминал состояния доски
		fmt.Print(" ")
		for i := range boardSize {
			fmt.Printf("%d", i+1) // выводим номера столбцов
		}
		fmt.Println()
		for i := range boardSize {
			fmt.Printf("%d", i+1) // выводим номера строк
			for j := range boardSize {
				switch board[i][j] {
				case empty:
					fmt.Print(" ")
				case cross:
					fmt.Print("X")
				case nought:
					fmt.Print("O")
				}
			}
			fmt.Println()
		}
		// Завершение вывода в терминал

		// Основной игровой цикл
		for state == playing {
			// Вывод сообщения о ходе текущего игрока
			playerSymbol := "X"
			if currentPlayer == nought {
				playerSymbol = "O"
			}

			fmt.Printf("%s's turn. Enter row and column (e.g., 1 2): ", playerSymbol)

			validInput := false
			for !validInput {
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input:", err)
					continue
				}

				input = strings.TrimSpace(input)
				if input == "q" {
					state = quit
					break
				}

				parts := strings.Fields(input)
				if len(parts) != 2 {
					fmt.Println("Invalid input. Please try again.")
					continue
				}

				row, err1 := strconv.Atoi(parts[0])
				col, err2 := strconv.Atoi(parts[1])

				if err1 != nil || err2 != nil {
					fmt.Println("Invalid input. Please try again.")
					continue
				}

				if row < 1 || row > boardSize || col < 1 || col > boardSize {
					fmt.Println("Invalid input. Please try again.")
					continue
				}
				// приведение к 0-индексации
				row--
				col--

				if board[row][col] != empty {
					fmt.Println("This cell is already occupied!")
					continue
				}

				// Выполнение хода
				board[row][col] = currentPlayer

				// Проверка выигрыша по строкам и столбцам
				winFound := false
				for i := range boardSize {
					rowWin := true
					colWin := true

					for j := range boardSize {
						if board[i][j] != currentPlayer {
							rowWin = false
						}
						if board[j][i] != currentPlayer {
							colWin = false
						}
					}

					if rowWin || colWin {
						winFound = true
						break
					}
				}

				// Проверка выигрыша по диагоналям
				if !winFound {
					diagWin := true
					for i := range boardSize {
						if board[i][i] != currentPlayer {
							diagWin = false
							break
						}
					}
					if diagWin {
						winFound = true
					}
				}

				// Проверка обратной диагонали
				if !winFound {
					antiDiagWin := true
					for i := range boardSize {
						if board[i][boardSize-i-1] != currentPlayer {
							antiDiagWin = false
							break
						}
					}
					if antiDiagWin {
						winFound = true
					}
				}

				if winFound {
					if currentPlayer == cross {
						state = crossWin
					} else {
						state = noughtWin
					}
				} else {
					// Проверка на ничью
					full := true
					for i := range boardSize {
						for j := range boardSize {
							if board[i][j] == empty {
								full = false
								break
							}
						}
						if !full {
							break
						}
					}
					if full {
						state = draw
					}
				}
				// Вывод состояния доски
				fmt.Print(" ")
				for i := range boardSize {
					fmt.Printf("%d", i+1) // выводим номера столбцов
				}
				fmt.Println()
				for i := range boardSize {
					fmt.Printf("%d", i+1)
					for j := range boardSize {
						switch board[i][j] {
						case empty:
							fmt.Print(". ")
						case cross:
							fmt.Print("X ")
						case nought:
							fmt.Print("O ")
						}
					}
					fmt.Println()
				}

				// Вывод сообщения о результате игры
				if state == crossWin {
					fmt.Println("Cross wins!")
				} else if state == noughtWin {
					fmt.Println("Nought wins!")
				} else if state == draw {
					fmt.Println("It's a draw!")
				} else {
					if currentPlayer == cross {
						currentPlayer = nought
					} else {
						currentPlayer = cross
					}
				}
				validInput = true
			}
		}
	}

}
