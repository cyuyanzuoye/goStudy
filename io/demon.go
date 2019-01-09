package main

import (
	"errors"
	"fmt"
	"io"
)

const (
	mutexLocked           = 1 << iota // mutex is locked  1<<0   1
	mutexWoken                        //2   1<<1
	mutexStarving                     //4   1<<2
	mutexWaiterShift      = iota      //3
	starvationThresholdNs = 1e6
)

func tt() {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	fmt.Println(starvationThresholdNs)
}

type test struct {
	data []int //操作的数据
	rix  int
	wix  int
}

func (t *test) Read(p []int) (n int, err error) {

	l := len(p)
	num := 0
	//右侧边界越界
	if t.rix > len(t.data) {
		return 0, nil
	}

	//边界裁切
	if (t.rix + l) < len(t.data) {
		num = l
	} else {
		num = len(t.data) - t.rix
	}

	//循环给p设置数据
	fmt.Println(num)
	fmt.Println(t.rix)

	for ix := 0; ix < num; ix++ {
		p[ix] = t.data[t.rix+ix]
	}
	//增加读取下标
	t.rix = t.rix + num
	return num, nil
}

//io包中定义了非常多的interface
//只要实现了接口中的方法
//那么io包中的导出方法就可以传入我们自定义的对象然后进行处理
//像什么文件数据，网络数据，数据库数据都可以统一操作接口

type MyReWr struct {
	data string //保存的数据
	rix  int    //指向当前数据读取的位置下标
	wix  int    //指向当前数据写入的位置下标
}

//MyReWr的构造函数
func MyReWrNew(s string) *MyReWr {
	return &MyReWr{s, 0, 0}
}

//读取数据到p中
func (m *MyReWr) Read(p []byte) (n int, err error) {
	//获取读取字节的长度
	l := len(p)
	num := 0
	//读-边界检查
	if m.rix >= len(m.data) {
		return 0, errors.New("EOF")
	}
	//转为字节处理
	tmp := []byte(m.data)

	//判断当前数据读取的下标
	if (m.rix + l) < len(m.data) {
		num = l
	} else {
		num = len(m.data) - m.rix
	}
	//循环给p设置数据
	for ix := 0; ix < num; ix++ {
		p[ix] = tmp[m.rix+ix]
	}
	//增加读取下标
	m.rix = m.rix + num
	return num, nil
}

func (m *MyReWr) WriteString(test string) (n int, a error) {

	fmt.Println(test)
	return 1, nil
}

//将p中数据写入
func (m *MyReWr) Write(p []byte) (n int, err error) {
	l := len(p)
	num := 0
	if m.wix >= len(m.data) {
		return 0, errors.New("EOF")
	}
	tmp := []byte(m.data)
	//判断当前数据写入的下标
	if (m.wix + l) < len(m.data) {
		num = l
	} else {
		num = len(m.data) - m.wix
	}
	//循环写入数据
	for ix := 0; ix < num; ix++ {
		tmp[m.wix+ix] = p[ix]
	}
	m.data = string(tmp)
	//增加写入下标
	m.wix = m.wix + num
	return num, nil
}

func main() {
	tt()

	test1 := &test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 90, 1, 2, 3, 4, 1, 1}, 0, 0}
	ttt := make([]int, 5)
	for {
		nn, err := test1.Read(ttt)
		if nn == 0 || err != nil {
			break
		}
		fmt.Println(nn, ttt[0:nn])
	}

	//我们自定义的一个结构，实现了Read和Write方法
	m := MyReWrNew("12345678910")
	p := make([]byte, 3)
	//循环读取数据
	for {
		n, _ := m.Read(p)
		if n == 0 {
			break
		}
		fmt.Println(n, string(p[0:n]))
	}
	//循环写入数据
	for {
		n, _ := m.Write([]byte("111"))
		if n == 0 {
			break
		}
		fmt.Println(n, m.data)
	}

	//MyReWr结构就可以使用如下方法
	m2 := MyReWrNew("666666666")
	m3 := MyReWrNew("999999999")

	tee := io.TeeReader(m2, m3)
	for {
		//循环从m2中读取数据到p中，再写入到m3中。
		n, _ := tee.Read(p)
		if n == 0 {
			break
		}
		fmt.Println(m2, m3)
	}

	//向m4中拷贝m5中的数据
	m4 := MyReWrNew("aaaaaaaaa")
	m5 := MyReWrNew("mmmmmmm")
	io.Copy(m4, m5)
	fmt.Println(m4, m5)

	//从m6中读取数据放入p2中
	m6 := MyReWrNew("abcdefghijklmo")
	p2 := make([]byte, len(m6.data))
	io.ReadFull(m6, p2)
	fmt.Println(string(p2))

	//向m7中写入字符串，如果实现了WriteString方法会直接调用。
	m7 := MyReWrNew("hello")
	io.WriteString(m7, "world123")
	m7.WriteString("1231")
	fmt.Println(m7)
}
