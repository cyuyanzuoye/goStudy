package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println(ioutil.WriteFile("io/ioutil/file/WriteFile.txt", []byte("hello xmge"), 0666))
}
