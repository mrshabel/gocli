package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// returns the count for the total lines, words, bytes(runes)
func main() {
	filenames := os.Args[1:]

	if len(filenames) == 0 {
		readNoArgs()
		return
	}

	totalLc, totalWc, totalBc := 0, 0, 0

	for _, filename := range filenames {
		isDir, err := isDirectory(filename)
		if err != nil {
			// log error and proceed to next file
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if isDir {
			// log error and skip current file
			fmt.Printf("%s: Is a directory\n", filename)
			continue
		}

		file, err := os.Open(filename)
		if err != nil {
			// log error and proceed to next file
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// initialize variables for metrics to be collected
		lc, wc, bc := 0, 0, 0

		// read file contents, line by line
		scan := bufio.NewScanner(file)

		// process input contents until there is none
		for scan.Scan() {
			text := scan.Text()
			// record word-count by splitting the line using any whitespace character such as space or tab
			wc += len(strings.Fields(text))
			// record byte-count as length of current token
			bc += len(text)

			// increment line count
			lc++
		}

		// neatly log the stats to the console. |lc		wc		bc|
		fmt.Printf("%3d %3d %3d %s\n", lc, wc, bc, filename)

		// add to aggregated totals
		totalLc += lc
		totalWc += wc
		totalBc += bc

		// close file at the end of current iteration to avoid maxing file descriptors
		file.Close()
	}

	// neatly log the total stats to the console. |lc		wc		bc|
	if len(filenames) > 0 {
		fmt.Printf("%3d %3d %3d total\n", totalLc, totalWc, totalBc)
	}
}

func readNoArgs() {
	// initialize variables for metrics to be collected
	lc, wc, bc := 0, 0, 0

	// read file contents, line by line
	scan := bufio.NewScanner(os.Stdin)

	// process input contents until there is none
	for scan.Scan() {
		text := scan.Text()
		// record word-count by splitting the line using any whitespace character such as space or tab
		wc += len(strings.Fields(text))
		// record byte-count as length of current token
		bc += len(text)

		// increment line count
		lc++
	}

	// neatly log the stats to the console. |lc		wc		bc|
	fmt.Printf("%3d %3d %3d\n", lc, wc, bc)
}

// read the contents of a given io reader (file, or stdin)
// func readContents(r io.Reader) {

// }

// check if a provided file path is a directory or not
func isDirectory(path string) (bool, error) {
	// get path stat
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}
