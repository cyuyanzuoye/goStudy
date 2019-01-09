package main

import (
	"encoding/json"
	"fmt"
)

func init() {
	fmt.Println("JSON学习过程--Map和JSon")
}

func main() {
	JsonToMap()

}

/**
  JSon的key值是字符串类型
*/
func JsonToMap() {
	test := `
解析到Map
我们知道interface{}可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的json数据的结果。
JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：
bool代表JSON booleans
float64代表JSON numbers
string代表JSON strings
nil 代表JSON null`
	fmt.Println(test)

	//jison解析
	jsonMap := make(map[string]interface{}, 10)
	jsonMap["1"] = "123"
	jsonMap["2"] = 1
	jsonMap["3"] = 1.22
	jsonMap["4"] = false
	jsonMap["5"] = nil
	jsonMap["6"] = []int{1, 2, 3, 4}
	jsonMap["6"] = map[int]interface{}{1: 2, 2: 3, 3: "a"} //key为字符串
	fmt.Println(jsonMap["1"])
	jsonStr, _ := json.MarshalIndent(jsonMap, "", "	")
	fmt.Println(string(jsonStr))

	//json字符串解析
	jsonMap2 := make(map[string]interface{}) //错误
	json.Unmarshal(jsonStr, &jsonMap2)       //错误
	fmt.Println(jsonMap2)

	//使用 value.(type)的方式根据不同的类型处理
	for key, value := range jsonMap2 { //只是点1-- value.(type) 只适合switch
		fmt.Println(key)
		fmt.Println(value)
	}

	//验证是否是Json字符串
	fmt.Println(json.Valid([]byte(`{"1":1}`))) //{"1":"1"} ture {1:1} 为false
	//f := interface{}

}
