package main

import (
	"flag"
	"fmt"
	"github.com/KEINOS/go-countline/cl"
	"log"
	"os"
)

var count = flag.Bool("c", false, "count bytes in given file")
var lines = flag.Bool("l", false, "count lines in given files")

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No file specified")
		os.Exit(1)
	}

	filename, fi, _ := getFileStats()

	if *count {
		fmt.Println(fi.Size(), filename)
	}

	if *lines {
		filename, _, file := getFileStats()

		count, err := cl.CountLines(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(count, filename)
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

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("Error retrieving statistics:", err)
		os.Exit(1)
	}
	return filename, fi, file
}
