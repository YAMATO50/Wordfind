package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func loadDatabase() {
	jsonData, err := ioutil.ReadFile("wordListDatabase.json")
	noDatabaseFile := errCheck(err, os.IsNotExist)
	if noDatabaseFile {
		return
	}

	err = json.Unmarshal(jsonData, &mainDatabase)
	errCheck(err, isNoErr)
}

func saveDatabase() {
	jsonData, err := json.Marshal(mainDatabase)
	errCheck(err, isNoErr)

	err = ioutil.WriteFile("wordListDatabase.json", jsonData, 0664)
	errCheck(err, isNoErr)
}
