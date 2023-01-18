package main

import (
	"bufio"
	"fmt"
	"math"
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
	text = strings.Replace(text, "\r", "", -1)
	return text
}

func guess(characters string, length int) []string {
	sameLengthWords := mainDatabase.WordLength[length]

	hashes := getPossibleHashes(characters)

	var possibleWords []string
	for _, hash := range hashes {
		possibleWords = append(possibleWords, sameLengthWords.ClassifiedWords[hash]...)
	}

	return possibleWords
}

func getPossibleHashes(characters string) []uint32 {
	maxHash := hashWord(characters)

	components := splitHash(maxHash)

	var possibleHashes []uint32
	for i := 0; i < len(components); i++ {
		binArray := toBinaryArray(uint32(i) + 1)
		var hash uint32
		for j := 0; j < len(binArray); j++ {
			hash = components[j] * binArray[j]
		}
		possibleHashes = append(possibleHashes, hash)
	}
	return possibleHashes
}

func splitHash(maxHash uint32) (components []uint32) {
	remainders := toBinaryArray(maxHash)

	for i, remainder := range remainders {
		component := uint32(math.Pow(2, float64(i))) * remainder
		if component == 0 {
			continue
		}
		components = append(components, component)
	}
	return
}

func toBinaryArray(num uint32) []uint32 {
	var remainders []uint32
	for num != 0 {
		remainders = append(remainders, num%2)
		num = num / 2
	}
	return remainders
}