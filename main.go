package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	Project Parameters:

	1. The program should take two arguments - the search string and a directory path. The task of the
		program is to search all text files within that directory (recursively, so it includes
		subdirectories) for the search string.

	2. The search in each file should be performed concurrently - each file should be processed by a
		separate goroutine.

	3. When a match is found, the program should print the file name and the line number where the
		match was found.

	4. Since the goroutines will be writing to the console concurrently, you'll need to ensure that
		their output doesn't interleave. You can use a Mutex for this purpose.

	5. Keep track of how many files have been processed and print this number once all files have
		been searched.
*/

func main() {
	term := os.Args[1]

	files, e := os.ReadDir("./testfiles")
	if e != nil {
		log.Fatal(e)
	}

	contents := make(map[string][][]byte, len(files))

	for _, f := range files {
		content, e := os.ReadFile(fmt.Sprintf("./testfiles/%s", f.Name()))
		if e != nil {
			log.Fatal(e)
		}

		contents[f.Name()] = Split(content, 10)
	}

	var hits []string

	for fname, lines := range contents {
		for i, line := range lines {
			if strings.Contains(string(line), term) {
				hits = append(hits, fmt.Sprintf("%s:%d", fname, i+1))
			}
		}
	}

	fmt.Println(hits)
}

func Split(input []byte, delimiter byte) [][]byte {
	var result [][]byte
	buffer := bytes.Buffer{}

	for _, b := range input {
		if b == delimiter {
			result = append(result, append([]byte(nil), buffer.Bytes()...))
			buffer.Reset()
		} else {
			buffer.WriteByte(b)
		}
	}

	if buffer.Len() > 0 {
		result = append(result, append([]byte(nil), buffer.Bytes()...))
	}

	return result
}

func FindIn(term string, contents [][]byte) string {
	return ""
}
