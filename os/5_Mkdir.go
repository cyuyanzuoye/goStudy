package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Chdir("20190105/os/file/")
	if err != nil {
		log.Fatal(err)
	}

	// Mkdir creates a new directory with the specified name and permission
	// bits (before umask).
	// If there is an error, it will be of type *PathError.
	fmt.Println(os.Mkdir("test_dir5", 777))
}
