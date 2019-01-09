/*
该软件包不支持磁盘跨越。

有关 ZIP64 的说明：

为了向后兼容，FileHeader 具有32位和64位大小字段。64位字段将始终包含正确的值，对于普通存档，这两个字段（32和64字段）都是相同的。对于需要 ZIP64 格式的文件，32位字段将为0xffffffff，必须使用64位字段替代。
 */


package main

import (
	"bytes"
	"os"
	"io"
	"log"
	"archive/zip"
	"path/filepath"
)

func isZip(zipPath string) bool {
	f, err := os.Open(zipPath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}
	//通过比对字节开头4个字节---PK\0x3\x04
	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

func zipFile(){
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	}{
		{"1.txt", "first"},
		{"2.txt", "second"},
		{"3.txt", "third"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("archive/zip/file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
}

func main() {
	zipFile()
	isZip("archive/zip/file.zip")
}
