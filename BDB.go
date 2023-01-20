package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func buildDatabase(newBuildDBfile string, newBuildDBfileExt string) {
	var wordList []string

	switch newBuildDBfileExt {
	case "txt":
		logActions("Loading txt file")
		wordList = txtToWordList(newBuildDBfile)
	case "list":
		logActions("Loading list")
		wordList = listToWordList(newBuildDBfile)
	case "dellist":
		deleteWordsFromDatabase(newBuildDBfile)
		saveDatabase()
		return
	}

	wordList = wordListToLower(wordList)

	logActions(fmt.Sprintf("%d words loaded", len(wordList)))
	logActions("Counting word lengths")
	wordLengthMap := countWordlength(wordList)

	logActions("Calculating hashes")
	preDatabase := characterizeByHash(wordLengthMap)

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

func listToWordList(list string) []string {
	words := strings.ReplaceAll(list, " ", "\n")
	return strings.Split(words, "\n")
}

func countWordlength(wordList []string) map[int][]string {
	wordLengthMap := make(map[int][]string)

	for _, word := range wordList {
		length := len(word)
		if length == 0 {
			continue //Ignore empty strings
		}

		wordsWithSameLength := wordLengthMap[length]
		if containsWord(wordsWithSameLength, word) {
			continue //Ignore doubled words
		}

		wordsWithSameLength = append(wordsWithSameLength, word)
		wordLengthMap[length] = wordsWithSameLength
	}
	return wordLengthMap
}

func containsWord(wordList []string, word string) bool {
	for _, element := range wordList {
		if element == word {
			return true
		}
	}
	return false
}

func characterizeByHash(wordLengthMap map[int][]string) Database {
	var db Database
	db.WordLength = make(map[int]SameLengthWord)

	for wordLength, wordList := range wordLengthMap {
		var slw SameLengthWord
		slw.ClassifiedWords = make(map[uint32][]string)

		for _, word := range wordList {

			hash := computeHash(word)

			sameHashWordList := slw.ClassifiedWords[hash]

			sameHashWordList = append(sameHashWordList, word)
			slw.ClassifiedWords[hash] = sameHashWordList
		}

		db.WordLength[wordLength] = slw
	}

	//Eliminating double words, because case is ignored
	for wordLength, slw := range db.WordLength {
		for hash, words := range slw.ClassifiedWords {
			for i := 0; i < len(words); i++ {
				for j := 0; j < len(words); j++ {
					if i == j {
						continue
					}
					if words[i] == words[j] {
						words[j] = words[len(words)-1]
						words = words[:len(words)-1]
					}
				}
			}
			slw.ClassifiedWords[hash] = words
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
			continue //no words with length 'length' contained in old database
		}

		mainDatabase.WordLength[length] = compareHashedWordMaps(mainSlw, preSlw)
	}

	return mainDatabase
}

var totalNewWords int

func compareHashedWordMaps(mainSlw SameLengthWord, preSlw SameLengthWord) SameLengthWord {
	for hash, sameHashWordList := range preSlw.ClassifiedWords {
		mainSameHashedWordList, ok := mainSlw.ClassifiedWords[hash]
		if !ok {
			mainSlw.ClassifiedWords[hash] = sameHashWordList
			continue //no words with hash 'hash' contained in old database
		}

		newWords := compareWordLists(mainSameHashedWordList, sameHashWordList)

		totalNewWords = totalNewWords + len(newWords) //needed for logging

		mainSameHashedWordList = append(mainSameHashedWordList, newWords...)
		mainSlw.ClassifiedWords[hash] = mainSameHashedWordList
	}
	return mainSlw
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

func wordListToLower(wordList []string) []string {
	for i := 0; i < len(wordList); i++ {
		wordList[i] = strings.ToLower(wordList[i])
	}
	return wordList
}

var deletedWords int

func deleteWordsFromDatabase(list string) {
	words := strings.ReplaceAll(list, " ", "\n")
	wordList := strings.Split(words, "\n")
	wordList = wordListToLower(wordList)

	for _, word := range wordList {
		length := len(word)
		hash := computeHash(word)

		sameLengthWords, ok := mainDatabase.WordLength[length]
		if !ok {
			//word not in Database
			continue
		}
		hashedWords, ok := sameLengthWords.ClassifiedWords[hash]
		if !ok {
			//word not in Database
			continue
		}
		for idx, classifiedWord := range hashedWords {
			if classifiedWord != word {
				continue
			}

			//Delete word at position idx
			hashedWords[idx] = hashedWords[len(hashedWords)-1]
			hashedWords = hashedWords[:len(hashedWords)-1]

			deletedWords += 1 //needed for logging

			break //word Found
		}
		if len(hashedWords) == 0 {
			delete(mainDatabase.WordLength[length].ClassifiedWords, hash)
			continue
		}
		mainDatabase.WordLength[length].ClassifiedWords[hash] = hashedWords
	}
	logActions(fmt.Sprintf("From %d words, %d were deleted from the database", len(wordList), deletedWords))
}
