package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// read input from command line. could be more than one file
	filenames := os.Args[1:]

	// read from command line if no input is provided
	if len(filenames) == 0 {
		readNoArgs()
		return
	}

	for _, filename := range filenames {
		// check if file is valid
		isDir, err := isDirectory(filename)
		if err != nil {
			// write error to std and proceed to next file
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if isDir {
			// write error to std and proceed to next file
			fmt.Printf("%s: Is a directory\n", filename)
			continue
		}

		// open file and read its contents
		file, err := os.Open(filename)
		if err != nil {
			// write error to stderr and proceed to next file
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		defer file.Close()

		// read file contents in chunks and write to a buffer. this is done to ensure that huge files are not loaded into memory at once
		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, "failed to read contents of file", err)
			// proceed to next file
			continue
		}
	}

}

func readNoArgs() {
	// create a scanner instance to read text from stdin
	scan := bufio.NewScanner(os.Stdin)

	// write contents to output by advancing to next token in each iteration
	for scan.Scan() {
		// get the text and write it to console
		fmt.Println(scan.Text())
	}
}

// utility function for checking if a given file path is a directory or not
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}
