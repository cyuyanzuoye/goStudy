package main

import (
	"fmt"
	"encoding/xml"
	"os"
)

func init() {
	fmt.Println("xml解析")
}

func main() {
	test()


}


//处理和Json基本类似
func test(){
	//只认识大写的 和Json一样，使用xml的标签处理
	type Study struct {
		Id string   `xml:"id"`
		Name string
	}
	liming:=Study{"1","22"}
	test ,_:=xml.Marshal(liming)
	fmt.Println(string(test))

	//解析
	var liming2 Study
	_=xml.Unmarshal(test,&liming2)

	fmt.Println(liming2.Id)

	//证明不是同一个对象
	liming2.Id="121"
	fmt.Println(liming.Id)
}
