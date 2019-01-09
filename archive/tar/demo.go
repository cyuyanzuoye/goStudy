package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)
//Package tar 实现对 tar 档案的访问。它的目的是去涵盖大部分的变体（variations），其中包括 GNU 和 BSD tar生成的包
func main() {
	// 创建一个缓冲区来写入我们的存档。
	buf := new(bytes.Buffer)

	// 创建一个新的tar存档。
	tw := tar.NewWriter(buf)

	// 将一些文件添加到存档中。
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		//档案文件头信息
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		//添加当代头部信息
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		//档案的内容
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
	// 确保在Close时检查错误。
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	//保存
	f, err := os.OpenFile("archive/tar/file.rar", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)

	// 打开tar档案以供阅读。
	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)

	// 迭代档案中的文件。
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// tar归档结束
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}

}