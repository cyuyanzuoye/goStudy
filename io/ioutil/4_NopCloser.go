package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Create("io/ioutil/file/NopCloser.txt")
	io.WriteString(f, "hello xmge111111111111111111111111111")
	f.Seek(0, io.SeekStart)

	// NopCloser returns a ReadCloser with a no-op Close method wrapping
	// the provided Reader r.
	r := ioutil.NopCloser(f)
	buffer := make([]byte, 24)
	fmt.Println(r.Read(buffer))
	fmt.Println(string(buffer))

}
