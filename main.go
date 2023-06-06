package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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
	maxNumberOfLines := 0

	for _, f := range files {
		content, e := os.ReadFile(fmt.Sprintf("./testfiles/%s", f.Name()))
		if e != nil {
			log.Fatal(e)
		}

		contents[f.Name()] = Split(content, 10)
		maxNumberOfLines += len(contents[f.Name()])
	}

	var hits []string

	chnl := make(chan hit, maxNumberOfLines)

	var wg sync.WaitGroup

	for fname, lines := range contents {
		wg.Add(1)
		go func(fname string, lines [][]byte) {
			FindIn(fname, term, lines, chnl)
			wg.Done()
		}(fname, lines)
	}

	go func() {
		wg.Wait()
		close(chnl)
	}()

	for v := range chnl {
		hits = append(hits, fmt.Sprintf("%s:%d", v.fname, v.line))
	}
}

type hit struct {
	fname string
	line  int
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

func FindIn(fname string, term string, contents [][]byte, hits chan hit) {
	for i, line := range contents {
		if strings.Contains(string(line), term) {
			hits <- hit{fname: fname, line: i + 1}
			return
		}
	}
}
