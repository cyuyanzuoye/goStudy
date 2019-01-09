package main

import (
	"fmt"
)

func main() {
	arrayInit()
}

//数组初始化
func arrayInit() {
	a := [12]int{11, 2, 3, 4, 1, 1}
	var b string
	var d string
	fmt.Scanf("%s", &d)
	fmt.Scanf("%s", &b)
	fmt.Printf("u", b)
	fmt.Printf("u", a)
	fmt.Printf("u", d)
}
