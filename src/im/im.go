package main

import (
	"time"
	"runtime"
	"flag"
	"fmt"
        "math/rand"
	"net"
	"log"
	"im/lib"

)

var (
	VERSION    string = "1.0.0"
	BUILD_TIME string
	GO_VERSION string
	GIT_COMMIT_ID string
	GIT_BRANCH string
)

func main() {
	fmt.Printf("Version:     %s\nBuilt:       %s\nGo version:  %s\nGit branch:  %s\nGit commit:  %s\n", VERSION, BUILD_TIME, GO_VERSION, GIT_BRANCH, GIT_COMMIT_ID)
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	//if len(flag.Args()) == 0 {
	//	fmt.Println("usage: im config")
	//	return
	//}
	//访问启动
	//监听客户端
	listenClient();

	//测试协议部分

	//lib.WriteHeader(12,1,1,1,1,os.Stdout)


}

func listenClient(){
	Listen(handleClient,9999)
}

//支持Http Https
func handleClient( conn net.Conn)  {
	log.Println("handle new connection")
	//创建客户端处理
	client := lib.NewClient(conn)
	//client := NewClient(conn)
	client.Run()
}

//监听连接
func Listen(f func(net.Conn), port int) {
	address := fmt.Sprintf("0.0.0.0:%d", port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("listen err:%s", err)
		return
	}
	tcpListener, ok := listen.(*net.TCPListener)
	if !ok {
		log.Println("listen err")
		return
	}

	//监听客户端连接
	for {
		client, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Println("accept err:%s", err)
			return
		}
		f(client)
	}
}
