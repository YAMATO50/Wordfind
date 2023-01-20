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
		fmt.Print("Wordfind>")
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

	timeMeasurement()
	possibleWords := getPossibleWords(characters, length)
	elapsed := timeMeasurement()

	fmt.Println("Results:")
	fmt.Println()
	for _, word := range possibleWords {
		fmt.Println(word)
	}
	fmt.Println()

	if verbose {
		fmt.Printf("%d words found in %d ms\n", len(possibleWords), elapsed)
		return
	}
	fmt.Printf("%d words found\n", len(possibleWords))
}

var reader *bufio.Reader

func readLine() string {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	return text
}

func getPossibleWords(characters string, length int) []string {
	sameLengthWords := mainDatabase.WordLength[length]

	hashes := getPossibleHashes(characters)

	var possibleWords []string
	for _, hash := range hashes {
		possibleWords = append(possibleWords, sameLengthWords.ClassifiedWords[hash]...)
	}

	possibleWords = removeImpossibleWords(characters, possibleWords)

	return possibleWords
}

func getPossibleHashes(characters string) []uint32 {
	maxHash := computeHash(characters)

	components := splitHash(maxHash)

	var possibleHashes []uint32
	for i := 0; i < int(math.Pow(2, float64(len(components))))-1; i++ {
		binArray := toBinaryArray(uint32(i) + 1)
		var hash uint32
		for j := 0; j < len(binArray); j++ {
			hash += components[j] * binArray[j]
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

func removeImpossibleWords(characters string, mabyPossibleWords []string) []string {
	countedAvailableCharacters := countCharacters(characters)
	var possibleWords []string

	for _, word := range mabyPossibleWords {
		countetCharacters := countCharacters(word)
		possible := true
		for character, count := range countetCharacters {
			if count > countedAvailableCharacters[character] {
				possible = false
				break
			}
		}
		if possible {
			possibleWords = append(possibleWords, word)
		}
	}

	return possibleWords
}

func countCharacters(word string) map[string]int {
	singleChars := strings.Split(word, "")
	count := make(map[string]int)
	for _, char := range singleChars {
		count[char] += 1
	}
	return count
}
