package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1, err := os.Create("io/io/file/MultiReader1.txt")
	if err != nil {
		log.Fatal(err)
	}
	f1.WriteString("11111")
	f1.Seek(0, io.SeekStart)

	f2, err := os.Create("io/io/file/MultiReader2.txt")
	if err != nil {
		log.Fatal(err)
	}
	f2.WriteString("2")
	f2.Seek(0, io.SeekStart)

	f3, err := os.Create("io/io/file/MultiReader3.txt")
	if err != nil {
		log.Fatal(err)
	}

	f3.WriteString("3")
	f3.Seek(0, io.SeekStart)

	// MultiReader returns a Reader that's the logical concatenation of
	// the provided input readers. They're read sequentially. Once all
	// inputs have returned EOF, Read will return EOF.  If any of the readers
	// return a non-nil, non-EOF error, Read will return that error.
	r := io.MultiReader(f1, f2, f3)
	buffer := make([]byte, 5) //由于公用一个buff的时候需要处理，根据返回的n，进行截取

	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))

	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))

	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))

	fmt.Println(r.Read(buffer))
}
