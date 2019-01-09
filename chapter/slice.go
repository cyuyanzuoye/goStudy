package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Printf("测试切片demo")
	defer fmt.Println("this is a defer")
}

func main() {
	a, b := 10, 20
	testCreate()
	defer func(x int) {
		fmt.Println("defer:", x, b)
	}(a)

	a += 10
	b += 100

	fmt.Println("a=%d,b=%d\n", a, b)

	args := os.Args
	fmt.Println("err:xxx ip port", args)
}

//构建-
func testCreate() {
	array := []int{1, 2, 3, 4, 5}
	testSlice := array[1:3]
	fmt.Println(testSlice[1])
}
