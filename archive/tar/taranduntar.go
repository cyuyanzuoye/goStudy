package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"strings"
	"fmt"
)

/**
  文件压缩过程 <表示写的方向
  压缩  dest<----gzip<----tar<------   src（create）
  解压  srcTar-->gzip---->tar------>   dest(create)
 */

func main()  {
	file1,_:=os.Open("archive/tar/file/1.txt");
	file2,_:=os.Open("archive/tar/file/2.txt");
	file3,_:=os.Open("archive/tar/file/3.txt");
	file4,_:=os.Open("archive/tar/file/untar");
	files := []*os.File{file1,file2,file3,file4}
	Compress(files,"archive/tar/file/compress.tar")
	DeCompress("archive/tar/file/compress.tar","archive/tar/file/decompress")
}

//压缩 使用gzip压缩成tar.gz
func Compress(files []*os.File, dest string) error {
	//创建目标文件
	d, _ := os.Create(dest)
	defer d.Close()
	//创建gzip压缩器
	gw := gzip.NewWriter(d)
	defer gw.Close()
	//创建tar文档压缩器
	tw := tar.NewWriter(gw)
	defer tw.Close()

	//文件压缩
	for _, file := range files {
		err := compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

//文件压缩
func compress(file *os.File, prefix string, tw *tar.Writer) error {
	//获取文件状态
	info, err := file.Stat()
	if err != nil {
		return err
	}
	//如果是目录
	if info.IsDir() {
		fmt.Println("目录结构"+info.Name())
		prefix = prefix + "/" + info.Name()
		fmt.Println("目录结构"+prefix)
		//获取目录结构
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		fmt.Println("目录结构"+":重新读取目录结构信息")
		//处理目录文件
		for _, fi := range fileInfos {
			fmt.Println("目录结构"+":打开的文件"+file.Name() + "/" + fi.Name())
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			//压缩文件
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		//压缩文件
		header, err := tar.FileInfoHeader(info, "")
		fmt.Println("压缩文件")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		//文件信息
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		//拷贝文件
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压 tar.gz
func DeCompress(tarFile, dest string) error {
	//打开tar文件
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	//创建gzip解压器
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()

	//创建tar解压器
	tr := tar.NewReader(gr)
	for {
		//解压
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		//目标地址
		filename := dest + hdr.Name
		fmt.Println("解压文件路径"+filename)
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		//拷贝
		io.Copy(file, tr)
	}
	return nil
}

//创建文件
func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}