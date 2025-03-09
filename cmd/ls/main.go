package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	paths := os.Args[1:]

	for _, path := range paths {
		// check if provided path is a directory
		isDir, err := isDirectory(path)
		if err != nil {
			// log error and proceed to next skip
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if !isDir {
			fmt.Println(path)
			continue
		}

		dirs, err := os.ReadDir(path)
		if err != nil {
			// log error and proceed to next skip
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// retrieve directory names and log it to the console
		dirNames := make([]string, len(dirs))
		for i, dir := range dirs {
			dirNames[i] = dir.Name()
		}

		// log with parent directory name only when 2 or more directories were received in Args
		output := strings.Join(dirNames, "  ")
		if len(paths) > 1 {
			fmt.Printf("%s: \n%s \n\n", path, output)
			continue
		}

		fmt.Println(output)
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
