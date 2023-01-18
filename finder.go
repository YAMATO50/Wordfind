package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findWords() {
	reader = bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("wordFinder>")
		input := readLine()
		switch strings.ToLower(input) {
		case "exit":
			return
		case "find":
			find()
		default:
			fmt.Println("Use \"exit\" to exit or \"find\" to find words")
		}
	}

}

func find() {
	fmt.Print("Enter available characters: ")
	characters := readLine()
	fmt.Print("Enter word lenght: ")
	lengthString := readLine()
	length, err := strconv.Atoi(lengthString)
	if err != nil {
		fmt.Printf("%s is not a number!\n", lengthString)
		return
	}
	if length <= 0 || length > len(characters) {
		fmt.Println("Invalid length!")
		return
	}

	possibleWords := guess(characters, length)
	fmt.Println("Results:")
	fmt.Println()
	for _, word := range possibleWords {
		fmt.Println(word)
	}
	fmt.Println()
	fmt.Printf("%d words found\n", len(possibleWords))
}

var reader *bufio.Reader

func readLine() string {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func guess(characters string, length int) []string {
	return []string{}
}
