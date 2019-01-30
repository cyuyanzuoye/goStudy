// Package client provides ...
package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type Client struct {
	Server *ChatServer
	Name   string
	Conn   net.Conn
	lock   *sync.RWMutex
	Rooms  map[string]*Room
	In     chan *Message       // use no bufferd channel
	Out    chan *Message       // use no bufferd channel
	Quit   chan struct{}
}

func (c *Client) Listen() {
	//fmt.Printf("New client: %s\n", c.Name)
	for msg := range c.Out {
		//判断消息的类型
		switch msg.Command {
		//退出消息
		case QUIT:
			// broadcast to all rooms
			c.lock.RLock()
			for _, r := range c.Rooms {
				r.In <- msg
			}
			c.lock.RUnlock()
			c.Quit <- struct{}{}
			return
		//进入消息
		case JOIN:
			name := msg.Receiver
			room := c.Server.GetRoom(name)
			c.lock.Lock()
			c.Rooms[name] = room
			c.lock.Unlock()
			room.In <- msg
		//默认处理
		default:
			c.lock.RLock()
			room := c.Rooms[msg.Receiver]
			c.lock.RUnlock()
			room.In <- msg
		}
	}
}

func (c *Client) Resp() {
	//创建读写buf
	buf := bufio.NewWriter(c.Conn)
	isClosed := false
	for {
		//现在一个channel
		select {
		//如果是消息进入
		case msg := <-c.In:
		        //已经关闭
			if isClosed {
				continue
			}
		        //写入buf
			_, err := buf.Write([]byte(fmt.Sprintf(
				"%s %s:%s\n",
				msg.Time.Format(time.RFC3339),
				msg.Sender.Name,
				msg.Content,
			)))
		        //网络是否错误
			if ne, ok := err.(net.Error); ok {
				if ne.Timeout() || ne.Temporary() {
					continue
				} else {
					isClosed = true
				}
			} else {
				isClosed = true
			}
		//如果是退出命令
		case <-c.Quit:
			if !isClosed {
				buf.Flush()
			}
			c.Conn.Close()
			close(c.Out)
			close(c.In)
			close(c.Quit)
			c = nil
			return
		}
	}
}

func (c *Client) Recv() {
	buf := bufio.NewReader(c.Conn)
	var msg *Message

	for {
		line, err := buf.ReadString('\n')

		if err != nil || len(line) == 0 {
			if err == io.EOF || len(line) == 0 {
				fmt.Println(c.Name, " Remote Closed")
				msg = &Message{c, "", QUIT, fmt.Sprintf("%s Lefted", c.Name), time.Now()}
			} else {
				log.Println(c.Conn.RemoteAddr(), "Error: ", err)
				msg = &Message{c, "", QUIT, fmt.Sprintf("%s DISCONNECT", c.Name), time.Now()}
			}
			//消息写入
			c.Out <- msg
			break
		} else {
			data := strings.Split(strings.TrimSpace(line), " ")
			if len(data) != 2 {
				continue
			}
			//切割消息
			room, content := data[0], data[1]
			msg = &Message{
				Sender:   c,
				Receiver: room,
				Content:  content,
				Time:     time.Now(),
			}

			//消息命令
			c.lock.RLock()
			if _, ok := c.Rooms[room]; ok {
				msg.Command = NORMAL
			} else {
				msg.Command = JOIN
			}
			c.lock.RUnlock()
		}
		//消息透传
		c.Out <- msg
	}
}
