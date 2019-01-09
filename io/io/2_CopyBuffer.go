package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1, err := os.Create("io/io/file/CopyBuffer1.txt")
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Create("io/io/file/CopyBuffer2.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(f1, "hello xmge")
	if err != nil {
		log.Fatal(err)
	}
	f1.Seek(0, os.SEEK_CUR)

	// CopyBuffer is identical to Copy except that it stages through the
	// provided buffer (if one is required) rather than allocating a
	// temporary one. If buf is nil, one is allocated; otherwise if it has
	// zero length, CopyBuffer panics.
	fmt.Println(io.CopyBuffer(f2, f1, make([]byte, 5)))
}
