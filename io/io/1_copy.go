package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f1, err := os.Create("io/io/file/Copy1.txt")
	if err != nil {
		log.Fatal(err)
	}
	//知识点1-go语言很神奇-必须存在目录才能创建目录，故这个方法需要先创建目录
	f2, err := os.Create("io/io/file/Copy2.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(f1, "你好啊11")
	if err != nil {
		log.Fatal(err)
	}
	//知识点2-由于对f1进行了些操作-但是没有移动f1的文件指针其指向末尾-故copy是不成功的
	f1.Seek(0, io.SeekStart)
	fmt.Println(io.Copy(f2, f1))
}
