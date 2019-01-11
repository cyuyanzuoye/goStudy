package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)
//Package tar 实现对 tar 档案的访问。它的目的是去涵盖大部分的变体（variations），其中包括 GNU 和 BSD tar生成的包
func main() {
	files:=[]string {
		"archive/tar/file/1.txt",
		"archive/tar/file/2.txt",
		"archive/tar/file/untar",
		"archive/tar/file/3.txt"}
	Tar(files, "archive/tar/file/test.tar")

	//UnTar("archive/tar/file/test.tar","archive/tar/file/untar")
}

func test(){
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

//添加
func Tar(files []string ,target string ){
	//创建目标文件
	f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//创建tar压缩器
	tw := tar.NewWriter(f)
	defer tw.Close()

	//临时添加文件
	var tempFile *os.File
	defer  tempFile.Close()
	
	//临时文件信息
	var tempFileInfo os.FileInfo
	//创建目标
	for _ ,fileName :=range files{
		tempFile,err =os.Open(fileName)
		if err!=nil{
			log.Fatal(err)
			continue;
		}
		tempFileInfo ,err=tempFile.Stat()

		//如果是目录-则添加目录
		if err!=nil{
			log.Fatal(err)
			continue;
		}

		//档案文件头信息
		hdr := &tar.Header{
			Name: tempFileInfo.Name(),
			Mode: 0600,
			Size: tempFileInfo.Size(),
		}
		//添加当代头部信息
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
			continue;
		}

		//档案的内容
		for{
			n,err:=io.Copy(tw,tempFile)
			if n==0 || err==io.EOF {
				break;
			}
			if err!=nil {
				log.Fatal(err)

			}

		}
	}

}

//解
func UnTar(src ,target string) (err error){

	//获取tar操作的文件
	f,err:=os.Open(src)
	if err!=nil {
		log.Fatal(err)
		return err
	}

	//创建解压目录
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	targetDir ,err:=os.Open(target)
	targetDir.Close()
	if err!=nil {
		return err
	}


	//创建解压器
	tr :=tar.NewReader(f)
	defer f.Close()

	var tempFile *os.File
	defer  tempFile.Close()

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
		fmt.Printf("Contents of %s:\n", hdr.FileInfo().IsDir())

		//目录文件-则创建目录
		path := filepath.Join(target, hdr.Name)
		if hdr.FileInfo().IsDir() {
			os.MkdirAll(path, hdr.FileInfo().Mode())
			continue
		}

		//创建文件
		tempFile,err := os.Create(path)
		if err!=nil {
			log.Fatal(err)
		}

		//解压 Copy 默认是3k的拷贝
		for  {
			n ,err :=io.Copy(tempFile,tr)
			if err==io.EOF || n==0 {
				break;
			}
			if err!=nil {
				log.Fatal(err)
				break;
			}
		}

	}
	return
}