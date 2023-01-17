package main

import (
	"fmt"
	"os"
	"strings"
)

var flagBuldDatabase bool
var newBuildDBfile string
var newBuildDBfileExt string
var verbose bool
var mainDatabase Database

func main() {
	loadDatabase()

	getFlags()
	if flagBuldDatabase {
		buildDatabase(newBuildDBfile, newBuildDBfileExt)
	}
}

func getFlags() {
	args := os.Args
	args = append(args, "-b")
	args = append(args, "wortliste.txt")
	for i, arg := range args {
		if strings.ToLower(arg) == "-b" {
			handleBuildDBFlag(args, i)
		}
		if strings.ToLower(arg) == "-v" {
			verbose = true
		}
	}
}

func handleBuildDBFlag(args []string, argPos int) {
	flagBuldDatabase = true

	if len(args) <= argPos+1 {
		fmt.Print("Missing filepath after -b flag")
		os.Exit(1)
	}

	newBuildDBfile = args[argPos+1]
	extFields := strings.Split(newBuildDBfile, ".")
	if len(extFields) == 1 {
		fmt.Println("Unsupported file extension")
		os.Exit(1)
	}

	newBuildDBfileExt = strings.ToLower(extFields[len(extFields)-1])
	switch newBuildDBfileExt {
	case "txt":
		break
	default:
		fmt.Println("Unsupported file extension")
		os.Exit(1)
	}
}

func logActions(logString string) {
	if verbose {
		fmt.Println(logString)
	}
}
