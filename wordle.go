package main

import (
	"fmt"
	"golang.org/x/term"
	"math/rand"
	"os"
	"strings"
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
			if current_letter > 0 {
				fmt.Println(current_letter)
				board[current_word-1][current_letter-1].Letter = " "
				if current_letter != 1 {
					current_letter--
				}
			}
		} else if unicode.IsLetter(r) { // Any Letter
			board[current_word-1][current_letter-1].Letter = string(r)
			if current_letter != 5 {
				current_letter++
			}
		} else if r == 27 { // Escape Key
			break
		} else if r == 13 { // Enter Key
			if board[current_word-1][4].Letter != " " {
				current_letter = 1
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

	fmt.Println(answer)
}

func print_board(board [6][5]Square) {
	for _, row := range board {
		fmt.Println("+-------+-------+-------+-------+-------+")
		fmt.Println("|       |       |       |       |       |")
		fmt.Printf("|   %v   |   %v   |   %v   |   %v   |   %v   |\n", row[0].Letter, row[1].Letter, row[2].Letter, row[3].Letter, row[4].Letter)
		fmt.Println("|       |       |       |       |       |")
	}
	fmt.Println("+-------+-------+-------+-------+-------+")
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
	randomNumber := rand.Intn(len(lines))

	return lines[randomNumber]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
