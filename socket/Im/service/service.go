package service

import (
	"fmt"
	"net"
	"strings"
)

const LISTEN_PORT = "9998"

func init() {
	fmt.Printf("servie-测试Im-服务器开发")
}

//服务器启动
func Start() {
	listener, err := net.Listen("tcp", LISTEN_PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			//链接中断--心跳检查
			continue
		}
		//TODO 达到携程上限-主动关闭
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	//此函数结束时，关闭连接套接字
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //n代码接收数据的长度
		if err != nil {
			fmt.Println(err)
			return
		}
		//切片截取，只截取有效数据
		result := buf[:n]
		//解析成request
		data := parse(result)
	}
}

//消息分发
func parse(data []byte) map[string]string {

}

//消息分发
func disthMessage() {

}
