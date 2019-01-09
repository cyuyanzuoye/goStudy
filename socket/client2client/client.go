package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

type groupClient struct {
	clients map[string]string
}

func NewGroupClient() *groupClient {
	return &groupClient{clients: make(map[string]string)}
}
func (g *groupClient) save(index string) {
	g.clients[index] = index
}
func (g *groupClient) remove(index string) {

	delete(g.clients, index)
}

func menu() {
	fmt.Println("-----------------------------------------")
	fmt.Println("1  退出房间                      ")
	fmt.Println("2  发送消息")
	fmt.Println("-----------------------------------------")
}

var groupClients *groupClient

func main() {
	//客户端主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err) //log.Fatal()会产生panic
		return
	}
	defer conn.Close() //关闭
	menu()

	groupClients = NewGroupClient()

	listenConn(conn)

}

//响应分发
func dispatachResponse(response map[string]string, conn net.Conn) (n int) {
	ipAddr := conn.RemoteAddr().String()
	addrArray := strings.Split(ipAddr, ":")
	clientIndex := addrArray[1]
	switch response["cmd"] {
	case "exit":
		//退出
		if clientIndex == response["client"] {
			fmt.Println("退出成功")
			return 1
		} else {
			groupClients.remove(response["client"])
		}

	case "send":
		//消息发送
		if response["code"] == "1" {
			fmt.Println("发送成功")
		} else {
			fmt.Println("消息发送失败")
		}

	case "login":
		//用户加入
		fmt.Printf(clientIndex)
		fmt.Printf(response["client"])
		groupClients.save(response["client"])
		if clientIndex != response["client"] {

		} else {
			fmt.Println("欢迎用户" + response["client"] + "加入聊天室")
		}
	case "tip":
		//提示
		fmt.Println(response["content"])
	default:
		//位置命令
		fmt.Println("发送未知命令")
	}
	return
}

func handleConn(conn net.Conn) {

	buf := make([]byte, 1024) //缓冲区
	var response map[string]string

	for {
		n, err := conn.Read(buf) //n代码接收数据的长度
		if err != nil {
			fmt.Println("已退出访问")
			return
		}
		fmt.Println(string(buf[:n]))
		//内容解析
		result := buf[:n]
		err = json.Unmarshal(result, &response)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if dispatachResponse(response, conn) == 1 {
			return
		}
	}
}

func listenConn(conn net.Conn) {
	var memuIndex int
	var to string
	var message string
	command := make(map[string]string)

	for {
		go handleConn(conn)

		fmt.Println("请选择输入命令")
		fmt.Scan(&memuIndex)
		switch memuIndex {
		case 1:
			command["cmd"] = "exit"
			commandJson, err := json.Marshal(command)
			if err != nil {
				fmt.Println("程序出现错误")
			}
			conn.Write([]byte(commandJson))
		case 2:
			fmt.Println(groupClients.clients)
			if len(groupClients.clients) <= 0 {
				fmt.Println("没有可发送的用户")
				break
			}

			command["cmd"] = "send"
			fmt.Println("请输入发送对象")
			for _, clientIndex := range groupClients.clients {
				fmt.Println("用户编号:" + clientIndex)
			}
			fmt.Scan(&to)
			command["to"] = to

			fmt.Println("请输入发送的消息")
			fmt.Scan(&message)
			command["content"] = message

			commandJson, err := json.Marshal(command)
			if err != nil {
				fmt.Println("程序出现错误")
			}
			conn.Write([]byte(commandJson))

		default:
			fmt.Println("其他命令展示不支持")
		}
	}

}
