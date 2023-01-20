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

		wordsWithSameLength = append(wordsWithSameLength, word)
		wordLengthMap[length] = wordsWithSameLength
	}
	return wordLengthMap
}

func characterizeByHash(wordLengthMap map[int][]string) Database {
	var db Database
	db.SameLengthWords = make(map[int]SameLengthWord)

	for wordLength, wordList := range wordLengthMap {
		var slw SameLengthWord
		slw.SameHashedWords = make(map[uint32][]string)

		for _, word := range wordList {

			hash := computeHash(word)

			sameHashWordList := slw.SameHashedWords[hash]

			sameHashWordList = append(sameHashWordList, word)
			slw.SameHashedWords[hash] = sameHashWordList
		}

		db.SameLengthWords[wordLength] = slw
	}

	//Eliminating double words, because case is ignored
	db = eliminateDoubleWords(db)

	return db
}

func eliminateDoubleWords(db Database) Database {
	for wordLength, slw := range db.SameLengthWords {
		for hash, words := range slw.SameHashedWords {

			for i := 0; i < len(words); i++ {
				for j := 0; j < len(words); j++ {

					if i == j {
						continue //Word at the same index is always the same
					}

					if words[i] == words[j] {
						words = deleteFromSlice(words, j)
					}
				}
			}
			slw.SameHashedWords[hash] = words
		}

		db.SameLengthWords[wordLength] = slw
	}

	return db
}

func deleteFromSlice(s []string, idx int) []string {
	s[idx] = s[len(s)-1]
	return s[:len(s)-1]
}

func compareDatabases(mainDatabase Database, preDatabase Database) Database {
	for length, preSlw := range preDatabase.SameLengthWords {

		mainSlw, ok := mainDatabase.SameLengthWords[length]
		if !ok {
			mainDatabase.SameLengthWords[length] = preSlw
			continue //no words with length 'length' contained in old database
		}

		mainDatabase.SameLengthWords[length] = compareHashedWordMaps(mainSlw, preSlw)
	}

	return mainDatabase
}

var totalNewWords int

func compareHashedWordMaps(mainSlw SameLengthWord, preSlw SameLengthWord) SameLengthWord {
	for hash, sameHashWordList := range preSlw.SameHashedWords {
		mainSameHashedWordList, ok := mainSlw.SameHashedWords[hash]
		if !ok {
			mainSlw.SameHashedWords[hash] = sameHashWordList
			continue //no words with hash 'hash' contained in old database
		}

		newWords := compareWordLists(mainSameHashedWordList, sameHashWordList)

		totalNewWords = totalNewWords + len(newWords) //needed for logging

		mainSameHashedWordList = append(mainSameHashedWordList, newWords...)
		mainSlw.SameHashedWords[hash] = mainSameHashedWordList
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

		sameLengthWords, ok := mainDatabase.SameLengthWords[length]
		if !ok {
			//word not in Database
			continue
		}
		hashedWords, ok := sameLengthWords.SameHashedWords[hash]
		if !ok {
			//word not in Database
			continue
		}
		for idx, hashedWordFromDatabase := range hashedWords {
			if hashedWordFromDatabase != word {
				continue
			}

			hashedWords = deleteFromSlice(hashedWords, idx)

			deletedWords += 1 //needed for logging

			break //word Found
		}
		if len(hashedWords) == 0 {
			delete(mainDatabase.SameLengthWords[length].SameHashedWords, hash)
			continue //No need to save an empty map field
		}
		mainDatabase.SameLengthWords[length].SameHashedWords[hash] = hashedWords
	}
	logActions(fmt.Sprintf("From %d words, %d were deleted from the database", len(wordList), deletedWords))
}
