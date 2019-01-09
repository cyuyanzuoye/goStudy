package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	testReset()
}

func testFlate(){
	// 字典是一串字节。 压缩一些输入数据时，
	// 压缩器将尝试用找到的匹配替换子串
	// 在字典里。因此，字典应该只包含子字符串
	// 预计会在实际的数据流中找到。
	const dict = `<?xml version="1.0"?>` + `<book>` + `<data>` + `<meta name="` + `" content="`

	// 要压缩的数据应该（但不是必需的）包含频繁的数据
	// 子字符串匹配字典中的字符串。
	const data = `<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>
`

	var b bytes.Buffer

	// 使用特制字典压缩数据。
	zw, err := flate.NewWriterDict(&b, flate.DefaultCompression, []byte(dict))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(zw, strings.NewReader(data)); err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("压缩后的内容")
	fmt.Println(b.String())
	// 解压缩器必须使用与压缩器相同的字典。
	// 否则，输入可能显示为损坏。
	fmt.Println("Decompressed output using the dictionary:")
	zr := flate.NewReaderDict(bytes.NewReader(b.Bytes()), []byte(dict))
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	// 使用'＃'替代字典中的所有字节以直观地显示
	// 演示使用预设词典的大致效果。
	fmt.Println("Substrings matched by the dictionary are marked with #:")
	hashDict := []byte(dict)
	for i := range hashDict {
		hashDict[i] = '#'
	}
	zr = flate.NewReaderDict(&b, hashDict)
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}
}
//Reset 可以用来丢弃当前的压缩器或解压缩器状态，并利用先前分配的内存快速重新初始化它们
//简单来说就是Reset从新指定一个解压或者压缩的 buf空间
func testReset(){
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.\n",
		"Concurrency is not parallelism.\n",
		"The bigger the interface, the weaker the abstraction.\n",
		"Documentation is for users.\n",
	}
	//字符串读取器
	var r strings.Reader
	//字节读取器
	var b bytes.Buffer
	buf := make([]byte, 32<<10)

	//创建压缩器（木有关联）
	zw, err := flate.NewWriter(nil, flate.DefaultCompression)
	if err != nil {
		log.Fatal(err)
	}
	//创建解压器（木有关联）
	zr := flate.NewReader(nil)

	for _, s := range proverbs {
		//读取器重置
		r.Reset(s)
		b.Reset()
		fmt.Println(b.Len())
		// 重置压缩器并从某些输入流编码。
		zw.Reset(&b)
		if _, err := io.CopyBuffer(zw, &r, buf); err != nil {
			log.Fatal(err)
		}
		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		// 重置解压缩器并解码为某个输出流。
		if err := zr.(flate.Resetter).Reset(&b, nil); err != nil {
			log.Fatal(err)
		}
		if _, err := io.CopyBuffer(os.Stdout, zr, buf); err != nil {
			log.Fatal(err)
		}
		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}
	}
}