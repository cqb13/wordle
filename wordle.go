package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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

	//current_word := 0

	print_board(board)

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
