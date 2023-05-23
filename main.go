package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
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
	//term := os.Args[1]

	var wg sync.WaitGroup
	var mutex sync.Mutex

	files, e := os.ReadDir("./testfiles")
	if e != nil {
		log.Fatal(e)
	}

	for _, f := range files {
		wg.Add(1)
		go func(fi fs.DirEntry) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()

			/*
				The code below would require a mutex because of the multiple calls to fmt.Println.

				fmt.Println(" __,  _, _,_ _, _ __,")
				fmt.Println(" |_  / \\ | | |\\ | | \\")
				fmt.Println(" |   \\ / | | | \\| |_/")
				fmt.Println(" ~    ~  `~' ~  ~ ~  ")
				fmt.Println("")

			*/
			fmt.Println(fi.Name())
		}(f) // always explicitly pass this argument if looping + generating goroutines
	}

	wg.Wait()
}
