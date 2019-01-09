package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	println(123)
	BookToJson()
	JsonToBook()
	testReflect()
}

//测试反射
func testReflect() {
	book := Book{Id: 1, Name: "123", TT: 11, isValide: false}
	s := reflect.TypeOf(&book).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		jsonTagValue := s.Field(i).Tag.Get("json")
		fmt.Println(jsonTagValue) //将tag输出出来
		jsonTagValueArray := strings.Split(jsonTagValue, ",")
		for value, _ := range jsonTagValueArray {
			fmt.Println(value)
		}
	}
}

/**
  引structTag的目的是支持解析不同的结构类型进行转换
  json-tag的标题--`json:"标签定义"`
  json-tag 使用方式  `json:"标签名"`  ,本质上是通过反射，获取tag的内容，然后根据tag的内容进行解析
  go语言的struct的Json映射支持的tag
  _:忽略
  omitempty:忽略空

*/
type Book struct {
	//1 omitempy，可以在序列化的时候忽略0值或者空值
	/**
		字段的tag是“-”，那么这个字段不会输出到JSON
	tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中的serverName
	tag中如果带有“omitempty”选项，那么如果该字段值为空，就不会输出到JSON串中
	如果字段类型是bool,string,int,int64等，而tag中带有“,string”选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
	*/
	Id       int    `json:"id,omitempy,string" bss:"111"`
	Name     string `json:"name"`
	TT       int    `json:"-"`
	isValide bool   `json:"isValide,int"`
}

/**
 * 1 多值返回，不能单值接受==》否则报错
 * 2 接收的Json字符串和struct字段转换，先与json tag 匹配，无则使用默认匹配，否则默认为0
 * 3
 */
func BookToJson() {
	book := Book{Name: "123", TT: 1, isValide: true}
	bookJson, error := json.Marshal(book)
	if error != nil {
		println("错误", error.Error())
	}
	println(string(bookJson))
}

func JsonToBook() {
	book := `{"id":"123","name":"123","TT":123,"isValide":1}`
	var booktest Book
	error := json.Unmarshal([]byte(book), &booktest)
	if error != nil {
		println("错误", error.Error())
	}
	println(booktest.Id, booktest.Name, booktest.TT, booktest.isValide)
}
