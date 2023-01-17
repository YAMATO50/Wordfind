package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func buildDatabase(newBuildDBfile string, newBuildDBfileExt string) {
	//err := os.Mkdir("wordDataBase", 664)
	//errChek(err, os.IsExist)

	var wordList []string

	switch newBuildDBfileExt {
	case "txt":
		logActions("Loading txt file")
		wordList = txtToWordList(newBuildDBfile)
	}

	logActions(fmt.Sprintf("%d words Loaded", len(wordList)))
	logActions("Counting Word lengths")
	wordLengthMap := countWordlength(wordList)

	logActions("Calculating hashes")
	preDatabase := characterizeBySpecialCharacters(wordLengthMap)

	logActions("Comparing old database to added words")
	mainDatabase = compareDatabases(mainDatabase, preDatabase)
	logActions(fmt.Sprintf("From %d words %d were added to old database", len(wordList), totalNewWords))

	saveDatabase()
}

func isNoErr(err error) bool {
	return false
}

// checks error and ignores it if isErrFunc returns true, otherwise panics
func errCheck(err error, isErrFunc func(error) bool) bool {
	if isErrFunc(err) {
		return true
	}
	if err != nil {
		log.Fatal(err)
	}
	return false
}

func txtToWordList(filename string) []string {
	file, err := os.Open(filename)
	errCheck(err, isNoErr)
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	errCheck(err, isNoErr)

	words := strings.ReplaceAll(string(fileContent), " ", "\n")
	return strings.Split(words, "\n")
}

func countWordlength(wordList []string) map[int][]string {
	wordLengthMap := make(map[int][]string)

	for _, word := range wordList {
		length := len(word)
		if length == 0 {
			continue
		}

		wordsWithSpecificLength := wordLengthMap[length]
		if containsWord(wordsWithSpecificLength, word) {
			continue
		}

		wordsWithSpecificLength = append(wordsWithSpecificLength, word)
		wordLengthMap[length] = wordsWithSpecificLength
	}
	return wordLengthMap
}

func containsWord(slice []string, word string) bool {
	for _, element := range slice {
		if element == word {
			return true
		}
	}
	return false
}

func characterizeBySpecialCharacters(wordLengthMap map[int][]string) Database {
	var db Database
	db.WordLength = make(map[int]SameLengthWord)

	for wordLength, wordList := range wordLengthMap {
		slw := db.WordLength[wordLength]

		for _, word := range wordList {

			letterIndex := hashWord(word)
			specialLetterWordList := slw.classifiedWords[letterIndex]

			specialLetterWordList = append(specialLetterWordList, word)
			slw.classifiedWords[letterIndex] = specialLetterWordList
		}

		db.WordLength[wordLength] = slw
	}
	return db
}

func compareDatabases(mainDatabase Database, preDatabase Database) Database {
	for length, preSlw := range preDatabase.WordLength {

		mainSlw, ok := mainDatabase.WordLength[length]
		if !ok {
			mainDatabase.WordLength[length] = preSlw
			continue
		}

		mainDatabase.WordLength[length] = compareClassifiedWordMaps(mainSlw, preSlw)
	}

	return mainDatabase
}

var totalNewWords int

func compareClassifiedWordMaps(mainSlw SameLengthWord, preSlw SameLengthWord) (newSlw SameLengthWord) {
	for hash, classifiedWordList := range preSlw.classifiedWords {
		mainClassifiedWords, ok := mainSlw.classifiedWords[hash]
		if !ok {
			mainSlw.classifiedWords[hash] = classifiedWordList
			continue
		}

		newWords := compareWordLists(mainClassifiedWords, classifiedWordList)
		totalNewWords = totalNewWords + len(newWords)
		mainClassifiedWords = append(mainClassifiedWords, newWords...)
		mainSlw.classifiedWords[hash] = mainClassifiedWords
	}
	return
}

func compareWordLists(mainWordList []string, preWordList []string) (newWordList []string) {
	for _, word := range preWordList {
		wordContained := false
		for _, mainWord := range mainWordList {
			if word == mainWord {
				wordContained = true
				break
			}
		}

		if wordContained {
			continue
		}

		newWordList = append(newWordList, word)
	}
	return
}
