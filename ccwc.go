package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/KEINOS/go-countline/cl"
	"log"
	"os"
	"strings"
)

var shouldCountBytes = flag.Bool("c", false, "count bytes in given file")
var shouldCountLines = flag.Bool("l", false, "count lines in given file")
var shouldCountWords = flag.Bool("w", false, "count words in given file")
var shouldCountCharacters = flag.Bool("m", false, "count characters in given file")

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No file specified")
		os.Exit(1)
	}

	filename, fileStats, file := getFileStats()

	if *shouldCountBytes {
		fmt.Println(fileStats.Size(), filename)
	}

	if *shouldCountLines {
		count, err := cl.CountLines(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(count, filename)
	}

	if *shouldCountWords {
		scanner := bufio.NewScanner(file)
		wordCount := 0
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			wordCount += len(words)
		}

		fmt.Println(wordCount, filename)
	}

	if *shouldCountCharacters {
		content, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// https://stackoverflow.com/questions/12668681/how-to-get-the-number-of-characters-in-a-string
		// https://github.com/golang/go/issues/24923
		fmt.Println(len([]rune(string(content))), filename)
	}
}

func getFileStats() (string, os.FileInfo, *os.File) {
	args := flag.Args()
	filename := args[0]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	fileStats, err := file.Stat()
	if err != nil {
		fmt.Println("Error retrieving statistics:", err)
		os.Exit(1)
	}
	return filename, fileStats, file
}
