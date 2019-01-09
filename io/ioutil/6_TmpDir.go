package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// TempDir creates a new temporary directory in the directory dir
	// with a name beginning with prefix and returns the path of the
	// new directory. If dir is the empty string, TempDir uses the
	// default directory for temporary files (see os.TempDir).
	// Multiple programs calling TempDir simultaneously
	// will not choose the same directory. It is the caller's responsibility
	// to remove the directory when no longer needed.
	name, err := ioutil.TempDir("io/ioutil/file/", "tmpDir-")
	if err != nil {
		log.Fatal(nil)
	}
	fmt.Println(name)
}
