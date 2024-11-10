package main

import (
	"fmt"
	"golang.org/x/term"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func main() {
	words := load_words()
	answer := pickWord(words)

	board := [6][5]string{
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " "},
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

		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			break
		}

		var r rune = rune(b[0])

		if r == 127 {
			if current_letter > 0 {
				fmt.Println(current_letter)
				board[current_word-1][current_letter-1] = " "
				if current_letter != 1 {
					current_letter--
				}
			}
		} else if unicode.IsLetter(r) {
			board[current_word-1][current_letter-1] = string(r)
			if current_letter != 5 {
				current_letter++
			}
		} else if r == 27 {
			break
		} else if r == 13 {
			if board[current_word-1][4] != " " {
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

func print_board(board [6][5]string) {
	for _, row := range board {
		fmt.Println("+-------+-------+-------+-------+-------+")
		fmt.Println("|       |       |       |       |       |")
		fmt.Printf("|   %v   |   %v   |   %v   |   %v   |   %v   |\n", row[0], row[1], row[2], row[3], row[4])
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
