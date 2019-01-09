package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("name", "1111")
	m.Store("stop", "2222")
	m.Store("sfsf", "3333")
	fmt.Println(m.Load("name"))
	fmt.Println(m.LoadOrStore("name", "aaa"))

	//// Range calls f sequentially for each key and value present in the map.
	//// If f returns false, range stops the iteration.
	m.Range(func(key, value interface{}) bool {
		fmt.Println("testsssss")
		fmt.Println(key, value)
		return false
	})
	fmt.Println(m.Load("name"))
}
