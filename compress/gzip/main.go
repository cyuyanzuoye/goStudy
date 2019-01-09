
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	//Example_writerReader()
	//ExampleReader_Multistream()

	CompressFile("compress/gzip/testdata/文章Compress.txt","compress/gzip/testdata/文章.txt")
	DeCompressFile("compress/gzip/testdata/文章Decompress.txt","compress/gzip/testdata/文章Compress.txt")
}

/**
   同样的--通过io.copy  ,将
 */
func Example_writerReader() {
	var buf bytes.Buffer




	//写入
	zw := gzip.NewWriter(&buf)

	// Setting the Header fields is optional.
	zw.Name = "a-new-hope.txt"
	zw.Comment = "an epic space opera by George Lucas"
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_,   err := zw.Write([]byte("A long time ago in a galaxy far, far away..."))



	if err != nil {
		log.Fatal(err)
	}

	n2 := buf.Len()

	if err := zw.Flush(); err != nil {
		log.Fatal(err)
	}

	n3 := buf.Len()
	if n2 == n3 {
		log.Fatal("Flush didn't flush any data")
	}


	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}



	//读入
	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}



	// Output:
	// Name: a-new-hope.txt
	// Comment: an epic space opera by George Lucas
	// ModTime: 1977-05-25 00:00:00 +0000 UTC
	//
	// A long time ago in a galaxy far, far away...
}


func ExampleReader_Multistream() {
	var buf bytes.Buffer           //操作的buf区域

	//吸入
	zw := gzip.NewWriter(&buf)

	var files = []struct {
		name    string
		comment string
		modTime time.Time
		data    string
	}{
		{"file-1.txt", "file-header-1", time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC), "Hello Gophers - 1"},
		{"file-2.txt", "file-header-2", time.Date(2007, time.March, 2, 4, 5, 6, 1, time.UTC), "Hello Gophers - 2"},
	}

	for _, file := range files {
		zw.Name = file.name
		zw.Comment = file.comment
		zw.ModTime = file.modTime

		if _, err := zw.Write([]byte(file.data)); err != nil {
			log.Fatal(err)
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		zw.Reset(&buf)
	}


	//读取
	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	for {
		zr.Multistream(false)
		fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}

		fmt.Print("\n\n")

		err = zr.Reset(&buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// Name: file-1.txt
	// Comment: file-header-1
	// ModTime: 2006-02-01 03:04:05 +0000 UTC
	//
	// Hello Gophers - 1
	//
	// Name: file-2.txt
	// Comment: file-header-2
	// ModTime: 2007-03-02 04:05:06 +0000 UTC
	//
	// Hello Gophers - 2
}

//压缩文件Src到Dst
func CompressFile(Dst string, Src string) error {
	//创建压缩后的文件
	newfile, err := os.Create(Dst)
	if err != nil {
		return err
	}
	defer newfile.Close()

	//打开待压缩的文件
	file, err := os.Open(Src)
	if err != nil {
		return err
	}

	//关联gzip写入对象
	zw := gzip.NewWriter(newfile)

	filestat, err := file.Stat()
	if err != nil {
		return nil
	}

	zw.Name = filestat.Name()
	zw.ModTime = filestat.ModTime()

	//待压缩文件，拷贝到gzip
	_, err = io.Copy(zw, file)
	if err != nil {
		return nil
	}

	//输出压缩
	//zw.Flush()
	if err := zw.Close(); err != nil {
		return nil
	}

	return nil
}

//解压文件Src到Dst----解压有问题
func DeCompressFile(Dst string, Src string) error {

	//打开解压缩文件
	file, err := os.Open(Src)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//创建解压文件
	newfile, err := os.Create(Dst)
	if err != nil {
		panic(err)
	}
	defer newfile.Close()
	file.Seek(0,io.SeekStart)


	buf :=make([]byte,1024*3)
	n,err:=file.Read(buf)

	bufdata :=bytes.NewBuffer(buf[:n])


	//关联gzip的读取路径--解压有问题
	zr, err := gzip.NewReader(bufdata)
	if err != nil  {
		panic(err)
	}

	filestat, err := file.Stat()
	if err != nil  {
		panic(err)
	}

	zr.Name = filestat.Name()
	zr.ModTime = filestat.ModTime()

	//执行拷贝操作
	_, err = io.Copy(newfile, zr)
	if err != nil &&err!=io.EOF{
		panic(err)
	}

	if err := zr.Close(); err != nil {
		panic(err)
	}
	return nil
}
