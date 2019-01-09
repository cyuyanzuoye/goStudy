package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1, err := os.Create("io/io/file/TeaReader1.txt")
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Create("io/io/file/TeaReader2.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(f1, "hello xmge")
	if err != nil {
		log.Fatal(err)
	}
	f1.Seek(0, io.SeekStart)

	// TeeReader returns a Reader that writes to w what it reads from r.
	// All reads from r performed through it are matched with
	// corresponding writes to w. There is no internal buffering -
	// the write must complete before the read completes.
	// Any error encountered while writing is reported as a read error.
	//TeaReader 是文件r读出，然后写入w
	r := io.TeeReader(f1, f2)
	buffer := make([]byte, 5)
	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))
}
