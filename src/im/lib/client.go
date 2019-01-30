package lib

import "net"


/**
  主要封装的是客户端的一些操作
 */
import (
	"fmt"
	"time"
)

const CLIENT_TIMEOUT  = 1000
type Client struct {
	Connection
}

func NewClient( conn net.Conn) *Client {
	client := new(Client)

	//初始化Connection
	client.conn = conn // conn is net.Conn or engineio.Conn

	//
	if netCoon, ok := conn.(net.Conn); ok {
		address := netCoon.LocalAddr()
		if taddr, ok := address.(*net.TCPAddr); ok {
			ip4 := taddr.IP.To4()
			publickIp := int32(ip4[0]) << 24 | int32(ip4[1]) << 16 | int32(ip4[2]) << 8 | int32(ip4[3])
			fmt.Println(publickIp)
		}
	}

	//TODO 创建读写通道
	//TODO 创建消息
	//TODO 通知消息
	return client

}

func (client *Client) Run()  {
	go client.Write()
	go client.Read()
}

//监听写入客户端数据
func (client *Client) Write(){
	fmt.Println("Write")
}

//监听读取客户端数据
func (client *Client) Read() {
	fmt.Println("Read")
	for {
		//原子操作，载入保存的数据
		if conn, ok := client.conn.(net.Conn); ok {
			//写入读取时间
			conn.SetReadDeadline(time.Now().Add(CLIENT_TIMEOUT * time.Second))
			msg := client.read()
			if msg == nil {
				client.HandleClientClosed()
				break
			}
			//消息处理
			client.HandleMessage(msg)
		}
	}
}

//处理客户端关闭的情况
func (client *Client) HandleClientClosed()  {
	client.close()
}

func (client *Client) HandleMessage(msg *Message){
	fmt.Println("msg cmd:",msg.cmd)                        //消息打印
	switch msg.cmd {
	//版本验证消息
	//case MSG_AUTH_TOKEN:
	//	client.HandleAuthToken(msg.body.(*AuthenticationToken), msg.version)
	//确认消息
	case MSG_ACK:
		client.HandleACK(msg.body.(*MessageACK))
	//心跳检查消息
	//case MSG_PING:
	//	client.HandlePing()
	}
	//处理-P2P消息， 组消息， 房间消息，自定义消息;
	//client.PeerClient.HandleMessage(msg)
	//client.GroupClient.HandleMessage(msg)
	//client.RoomClient.HandleMessage(msg)
	//client.CustomerClient.HandleMessage(msg)
}


func (client *Client) HandleACK(ack *MessageACK) {
	fmt.Println("ack:", ack.seq)
}

