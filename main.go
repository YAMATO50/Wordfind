package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var flagBuldDatabase bool
var newBuildDBfile string
var newBuildDBfileExt string
var verbose bool
var help bool
var mainDatabase Database

func main() {
	getFlags()
	if help {
		showHelp()
		return
	}
	loadDatabase()
	if flagBuldDatabase {
		logActions("Building database")
		timeMeasurement()
		buildDatabase(newBuildDBfile, newBuildDBfileExt)
		elapsed := timeMeasurement()
		logActions(fmt.Sprintf("Database building took %d ns", elapsed))
		return
	}
	findWords()
}

func getFlags() {
	args := os.Args
	for i, arg := range args {
		arg = strings.ToLower(arg)
		switch arg {
		case "-b":
			handleBuildDBFlag(args, i)
		case "-s":
			handleSingleWordFlag(args, i)
		case "-v":
			verbose = true
		case "-d":
			handleDeleteWordsFlag(args, i)
		case "-h":
			fallthrough
		case "-help":
			help = true
		default:
			continue
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

func handleSingleWordFlag(args []string, argPos int) {
	flagBuldDatabase = true

	if len(args) <= argPos+1 {
		fmt.Println("Missing at least one word after -s flag")
		os.Exit(1)
	}

	newBuildDBfile = strings.Join(args[argPos+1:], "\n")
	newBuildDBfileExt = "list"
}

func handleDeleteWordsFlag(args []string, argPos int) {
	flagBuldDatabase = true

	if len(args) <= argPos+1 {
		fmt.Println("Missing at least one word after -d flag")
		os.Exit(1)
	}

	newBuildDBfile = strings.Join(args[argPos+1:], "\n")
	newBuildDBfileExt = "dellist"
}

func logActions(logString string) {
	if verbose {
		fmt.Println(logString)
	}
}

var firstCall bool = true
var startTime int64

func timeMeasurement() int64 {
	if !verbose {
		return 0
	}

	if !firstCall {
		startTime = time.Now().UnixNano()
		firstCall = false
		return 0
	}

	firstCall = true
	return time.Now().UnixNano() - startTime
}

func showHelp() {
	fmt.Println("Usage: wordfind [flags|values]")
	fmt.Println()
	fmt.Println("Available flags:")
	fmt.Println("\t-h")
	fmt.Println("\t-help\tShows this help")
	fmt.Println("\t-b\tBuild/Update the database from the file specified in the argument after -b")
	fmt.Println("\t-s\tBuild/Update the database with all words following the -s flag")
	fmt.Println("\t-v\tVerbose output")
	fmt.Println("\t-d\tDelete all words following the -d flag from the Database")
	fmt.Println()
	fmt.Println("Available file Formats (-b)")
	fmt.Println("\t.txt\tWords separated by newline and/or whitespace")
}
