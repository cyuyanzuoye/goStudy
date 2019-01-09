package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type Group struct {
	id         int                 //组编号
	name       string              //组名
	clientList map[string]net.Conn //同组的客户端
	max        int                 //最大连接数
}

//创建实例
func NewGroup(id int, name string, max int) *Group {
	clientList := make(map[string]net.Conn)
	return &Group{id: id, name: name, max: max, clientList: clientList}
}

//组播
func (g *Group) Broad(message []byte) (n int) {
	//消息为空
	if len(message) <= 0 {
		return
	}
	if len(g.clientList) <= 0 {
		return
	}
	for _, conn := range g.clientList {
		n, err := conn.Write(message)
		if err != nil {

		}
		if n > 0 {
			n++
		}
	}
	return
}

//单个消息发送
func (g *Group) Send(message []byte, to string) (n int, err error) {
	if len(message) <= 0 {
		return 0, nil
	}
	if g.IsOnline(to) == false {
		return 0, errors.New("用户不在线")
	}
	return g.clientList[to].Write(message)
}

func (g *Group) IsOnline(index string) (status bool) {

	if len(g.clientList) < 0 {
		status = false
	}

	if g.clientList[index] == nil {
		status = false
	} else {
		status = true
	}
	return status
}

//上线
func (g *Group) OnLine(conn net.Conn, index string) (n int) {
	if g.IsOnline(index) {
		return
	}
	if (len(g.clientList) + 1) > g.max {
		return
	}

	g.clientList[index] = conn

	return 1
}

//下线
func (g *Group) OffLine(index string) (n int) {
	if !g.IsOnline(index) {
		return
	}

	delete(g.clientList, index)
	return 1
}

func init() {
	fmt.Println("聊天组服务器启动")
	fmt.Println("等待用户加入")
}

//连接对象处理
func dealConn(conn net.Conn) {
	response := make(map[string]string)
	defer conn.Close() //此函数结束时，关闭连接套接字

	//获取远程连接对象ip地址
	ipAddr := conn.RemoteAddr().String()

	//操作用户组
	if len(talkGroup.clientList) >= talkGroup.max {
		response["cmd"] = "tip"
		response["content"] = strings.ToUpper(string(talkGroup.name + ":抱歉，服务器拥堵"))
		data, _ := json.Marshal(response)
		conn.Write(data)
		return
	}

	data := strings.Split(ipAddr, ":")
	connIndex := data[1]

	//上线成功
	if talkGroup.OnLine(conn, connIndex) == 1 {
		response["cmd"] = "login"
		response["client"] = connIndex
		response["content"] = talkGroup.name + "：用户" + connIndex + "进入" + talkGroup.name + "聊天室"
		data, _ := json.Marshal(response)
		talkGroup.Broad(data)

		fmt.Println("用户" + connIndex + "进入" + talkGroup.name + "聊天室")
	}

	//缓冲区，用于接收客户端发送的数据
	buf := make([]byte, 1024)
	var commandMap map[string]string
	for {
		//阻塞等待用户发送的数据
		n, err := conn.Read(buf) //n代码接收数据的长度
		if err != nil {
			//退出命令
			response["cmd"] = "exit"
			response["client"] = connIndex
			response["content"] = "用户" + connIndex + "退出聊天室"
			data, _ := json.Marshal(response)
			talkGroup.Broad(data)

			talkGroup.OffLine(connIndex)
			fmt.Println("用户" + connIndex + "退出聊天室")

			return
		}

		//消息格式 A:你好涉及
		result := buf[:n]
		fmt.Printf("接收到数据来自[%s]==>[%d]:%s\n", ipAddr, n, string(result))

		//文本命令系统-JSON 交互 comand:内容:对象  send:{}
		err = json.Unmarshal(result, &commandMap)
		if err != nil {
			fmt.Println("内容解析失败")
			continue
		}

		//命令分发
		code := dispatachCmd(commandMap, conn, connIndex)
		if code == 1001 {
			fmt.Println("用户" + connIndex + "退出聊天室")
			return
		}
	}
}

//命令
func dispatachCmd(cmdMap map[string]string, conn net.Conn, connIndex string) (code int) {

	cmd := make(map[string]string)
	if cmd == nil {
		return
	}

	switch cmdMap["cmd"] {
	case "exit":
		//退出痕迹消除
		talkGroup.OffLine(connIndex)

		//发送广告命令
		cmd["cmd"] = "exit"
		cmd["client"] = connIndex
		response, err := json.Marshal(cmd)
		if err != nil {
			fmt.Println("用户" + connIndex + "要求退出失败")
		}
		talkGroup.Broad(response)

		return 1001
	case "send":
		//返回消息发送应答
		cmd["cmd"] = "send"
		cmd["client"] = cmdMap["to"]
		cmd["code"] = "1"
		cmd["msg"] = "发送成功"

		//获取发送内容
		_, err := talkGroup.Send([]byte(cmdMap["content"]), cmdMap["to"])
		if err != nil {
			cmd["cmd"] = "send"
			cmd["client"] = cmdMap["to"]
			cmd["code"] = "0"
			cmd["msg"] = "消息发送失败"

			talkGroup.OffLine(cmdMap["to"])
			fmt.Println("消息发送失败")

		}

		response, _ := json.Marshal(cmd)
		conn.Write(response)

		return 1002
	default:
		//发送广告命令
		cmd["cmd"] = "hehe"
		cmd["client"] = connIndex
		response, err := json.Marshal(cmd)
		if err == nil {
			conn.Write(response)
		}
		fmt.Println("未知的命令")
	}

	return 0
}

var talkGroup *Group

func main() {
	//创建、监听socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err) //log.Fatal()会产生panic
	}

	talkGroup = NewGroup(1, "掌阅", 5)

	defer listener.Close()

	//服务器-监听
	for {
		conn, err := listener.Accept() //阻塞等待客户端连接
		if err != nil {
			log.Println(err)
			continue
		}

		go dealConn(conn)
	}
}
