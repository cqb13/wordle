package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	words := load_words()
	word := pickWord(words)

	fmt.Println(word)
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
