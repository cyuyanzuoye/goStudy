package main

import (
	"fmt"
	"im/lib"
	"net"
	"log"
)

func init() {
	fmt.Println("protocolTest 正在启动")
}
func main() {
	//监听对抗
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
	       log.Fatal(err) //log.Fatal()会产生panic
		return
	}
	//消息通道
	channel := make(chan string)

	defer conn.Close()
	defer close(channel)

	menuID := 0;

	//
	go listenWrite(channel,conn)

	for{
		fmt.Println("*****************************************")
		fmt.Println("选项1:消息发送                           ")
		fmt.Println("*****************************************")
		fmt.Println("")
		fmt.Scanf("%d",&menuID);
		switch menuID {
		case 1:
			var messageStr string
			fmt.Println("消息内容")
			fmt.Scanf("%s",&messageStr)
			if(len(messageStr)>0){
				channel<- messageStr;
			}
		default:
			break;
		}

	}

}

func writeMessage(conn net.Conn,body []byte)  {
	lib.WriteHeader(int32(len(body)), int32(1), byte(1), byte(1), byte(1), conn)
	conn.Write(body)
}

func listenWrite(channel chan string ,conn net.Conn){
	for {
		messageStr,ok:= <-channel
		fmt.Println(messageStr)
		if !ok{
			fmt.Println(messageStr)
			log.Fatal(ok)
		}

		if(len(messageStr)>0){
			writeMessage(conn,[]byte(messageStr))
		}
	}
}
