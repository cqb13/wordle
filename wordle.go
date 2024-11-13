package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"slices"
	"strings"
	"time"
	"unicode"
)

type Status int

const (
	Correct Status = iota
	SemiCorrect
	Incorrect
)

type Square struct {
	Letter string
	Status Status
}

func NewSquare() Square {
	return Square{
		Letter: " ",
		Status: Incorrect,
	}
}

func main() {
	words := load_words()
	answer := pickWord(words)

	board := [6][5]Square{}

	for i := range board {
		for j := range board[i] {
			board[i][j] = NewSquare()
		}
	}

	clear_terminal()
	print_board(board)

	current_word := 1
	current_letter := 1

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error setting terminal to raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var b []byte = make([]byte, 1)

	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			break
		}

		term.Restore(int(os.Stdin.Fd()), oldState)

		var r rune = rune(b[0])

		if r == 127 { // Backspace Key
			board[current_word-1][current_letter-1].Letter = " "
			current_letter--
			if current_letter < 1 {
				current_letter = 1
			}
		} else if unicode.IsLetter(r) { // Any Letter
			board[current_word-1][current_letter-1].Letter = string(r)
			if current_letter < 5 {
				current_letter++
			}
		} else if r == 27 { // Escape Key
			break
		} else if r == 13 { // Enter Key
			if board[current_word-1][4].Letter != " " {
				// make sure the word is in the word list
				board[current_word-1] = check_accuracy(board[current_word-1], answer)
				player_win := check_player_win(board[current_word-1])
				if player_win {
					clear_terminal()
					print_board(board)
					fmt.Println("You got it!")
					break
				}
				current_letter = 1
				if current_word == 6 {
					clear_terminal()
					print_board(board)
					fmt.Println("You lost!")
					fmt.Printf("The word was %v\n", answer)
					break
				}
				current_word++
			}
		}

		clear_terminal()
		print_board(board)

		oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error re-setting terminal to raw mode:", err)
			return
		}
	}
}

func check_player_win(row [5]Square) bool {
	for _, square := range row {
		if square.Status != Correct {
			return false
		}
	}

	return true
}

func check_accuracy(row [5]Square, answer string) [5]Square {
	answer_letters := strings.Split(answer, "")

	for index, square := range row {
		if square.Letter == answer_letters[index] {
			row[index].Status = Correct
			answer_letters[index] = "-"
			continue
		}
	}

	for index, square := range row {
		if square.Status == Correct {
			continue
		}

		if slices.Contains(answer_letters, square.Letter) {
			for i, letter := range answer_letters {
				if letter == square.Letter {
					answer_letters[i] = "-"
					row[index].Status = SemiCorrect
					break
				}
			}
		}
	}

	return row
}

func print_board(board [6][5]Square) {
	for _, row := range board {
		fmt.Println("+-------+-------+-------+-------+-------+")
		fmt.Println("|       |       |       |       |       |")
		fmt.Printf("|   %v   |   %v   |   %v   |   %v   |   %v   |\n", make_colored(row[0]), make_colored(row[1]), make_colored(row[2]), make_colored(row[3]), make_colored(row[4]))
		fmt.Println("|       |       |       |       |       |")
	}
	fmt.Println("+-------+-------+-------+-------+-------+")
}

func make_colored(square Square) string {
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Reset = "\033[0m"

	if square.Status == Correct {
		return Green + square.Letter + Reset
	} else if square.Status == SemiCorrect {
		return Yellow + square.Letter + Reset
	} else {
		return square.Letter
	}
}

func clear_terminal() {
	esc := 27
	fmt.Printf("%c[2J%c[1;1H", esc, esc)
}

func load_words() []string {
	data, err := os.ReadFile("./words.txt")

	check(err)

	strData := string(data)

	lines := strings.Split(strData, "\n")

	return lines
}

func pickWord(lines []string) string {
	startDate := time.Date(2024, 11, 13, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now()
	duration := currentDate.Sub(startDate)

	days := int(duration.Hours()/24 + 0.5)
	if days > len(lines) {
		panic("Out of words!")
	}
	return lines[days]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
