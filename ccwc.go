package main

import (
	"flag"
	"fmt"
	"os"
)

var count = flag.Bool("c", false, "count bytes in given file")

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No file specified")
		os.Exit(1)
	}

	if *count {
		filename, fi := getFileStats()

		fmt.Println(fi.Size(), filename)
	}
}

func getFileStats() (string, os.FileInfo) {
	args := flag.Args()
	filename := args[0]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("Error retrieving statistics:", err)
		os.Exit(1)
	}
	return filename, fi
}
