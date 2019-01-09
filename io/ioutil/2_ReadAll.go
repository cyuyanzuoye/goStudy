package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileName := "io/ioutil/file/ReadAll.txt"
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("hello xmge11111111111111111111111111")
	f.Seek(0, io.SeekStart)

	// ReadAll reads from r until an error or EOF and returns the data it read.
	// A successful call returns err == nil, not err == EOF. Because ReadAll is
	// defined to read from src until EOF, it does not treat an EOF from Read
	// as an error to be reported.
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}
