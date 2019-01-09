package main

import (
	"encoding/json"
	"fmt"
)

func init() {
	fmt.Println("JSON学习结构体和JSON转换")
}

/**
默认输出的是 结构体的大写熟悉
*/
type Student struct {
	Age  int `json:"age,string"`
	Name int `json:"name,string"` // '知识点1-不能使用同类型标签,复杂类型'

}

func main() {
	student := Student{1, 1}
	fmt.Println(student.Age)

	//解析成字符串
	data, _ := json.Marshal(student)
	fmt.Println(string(data))

	//解析成结构体
	var animals Student
	err := json.Unmarshal([]byte(`{"age":"12","Name":"1"}`), &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
