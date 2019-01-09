package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Create("./log/log/log.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	flog := log.New(f, "[Debug] ", log.Llongfile)
	flog.Println("今天天气很好")
	flog.Fatal("1231231")
	flog.Panic("1231")
}
