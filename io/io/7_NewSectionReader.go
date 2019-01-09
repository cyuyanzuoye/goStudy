package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1, err := os.Create("io/io/file/NewSectionReader.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(f1, "hello xmge111")
	if err != nil {
		log.Fatal(err)
	}
	f1.Seek(0, io.SeekStart)

	// NewSectionReader returns a SectionReader that reads from r
	// starting at offset off and stops with EOF after n bytes.
	// 返回操作对象的一部分信息，不足，则返回部分。但是报错EOF
	r := io.NewSectionReader(f1, 5, 24)

	buffer := make([]byte, 1024)
	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))
}
