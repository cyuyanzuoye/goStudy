package main

import (
	"fmt"
	"math"
)

func main() {
	testJ()
}

func testAbs(){
	fmt.Println("%f",math.Abs(-1111.1))
	fmt.Printf("%f",math.Abs(-1111.1))
	fmt.Printf("%f",math.Abs(-123.2))
}

func testJ(){
	n:=0
	f:=1.2
	fmt.Println(math.J0(f))
	fmt.Println(math.J1(f))
	fmt.Println(math.Jn(n,f))
}

func testY(){

}
