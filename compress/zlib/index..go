package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func main(){
	ExampleNewReader();
	ExampleNewWriter();
}
/* 1  zlib 压缩使用-NewWriter 关联buf,通过zlib写入的内容，会压缩字节。转化成string是乱码，
         writestring ---->zlib-------->buf     通过操作buf 操作压缩后的字符
    2 zlib  解压缩使用 NewReader 关联buf
         buf------------->zlib-------->io.read------->abuf          通过操作abuf 操作解压字符
*/
func ExampleNewWriter() {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world你好是\n"))
	w.Close()
	b.String()
	fmt.Println(len("hello, world你好是\n"))
	fmt.Println(len(b.Bytes()))
	fmt.Println(b.String())
	// Output: [120 156 202 72 205 201 201 215 81 40 207 47 202 73 225 2 4 0 0 255 255 33 231 4 147]
	fmt.Println("执行ExampleNewWrite结束")
}

func ExampleNewReader() {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	b := bytes.NewReader(buff)

	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)
	// Output: hello, world
	r.Close()
	fmt.Println("执行ExampleNewReader结束")
}