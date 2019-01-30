package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)


func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("chat-room-client-start")
}

var (
	speed int64 = 0
)

func work(ch chan struct{}, idx int) {
	defer func() {
		<-ch
	}()
	//远程连接服务器
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Println(err)
		return
	}
	//创建bufIo
	buf := bufio.NewWriter(conn)

	//信息拷贝
	go func() {
		io.Copy(ioutil.Discard, conn)
	}()

	for {
		atomic.AddInt64(&speed, 1)
		time.Sleep(200 * time.Millisecond)
		//计算房间号
		room := strconv.Itoa(((idx-1)/100+1)*10 - 9 + rand.Int()&10)
		msg := room + " " + strconv.Itoa(rand.Int()) + "\n"

		//发送消息
		_, err := buf.Write([]byte(msg))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	go func() {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			fmt.Println(atomic.LoadInt64(&speed))
			atomic.StoreInt64(&speed, 0)
		}
	}()
	idx := 1
	ch := make(chan struct{}, 40000)
	for {
		ch <- struct{}{}
		go work(ch, idx)
		time.Sleep(2 * time.Millisecond)
		idx++
	}
}
