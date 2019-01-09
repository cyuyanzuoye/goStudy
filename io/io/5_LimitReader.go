package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1, err := os.Create("io/io/file/LimitReader.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(f1, "hello xmge11111111111111111111111111111111111111111111111111111111"+
		"11111111111111111111111111111111111111"+
		"11111111111111111111111111111111111111"+
		"1111111111111111111111111"+
		"111111111111111111111")
	if err != nil {
		log.Fatal(err)
	}

	// LimitReader returns a Reader that reads from r
	// but stops with EOF after n bytes.
	// The underlying implementation is a *LimitedReader.
	r := io.LimitReader(f1, 10)
	f1.Seek(0, os.SEEK_SET)

	buff := make([]byte, 20)
	n, err := r.Read(buff)
	if err != nil {
		fmt.Println(n)
	}
	fmt.Println(string(buff))

}
