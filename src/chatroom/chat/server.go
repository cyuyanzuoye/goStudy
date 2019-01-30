// Package server provides ...
package chat

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

//创建服务
func NewChatServer(bind string, rooms map[string]*Room, lock *sync.RWMutex) *ChatServer {
	return &ChatServer{bind, rooms, lock}
}

//聊天服务结构
type ChatServer struct {
	Bind_to string                          //绑定的端口号
	Rooms   map[string]*Room                //房间
	lock    *sync.RWMutex                   //信号量
}

// GetRoom return a room, if this name of room is not exist,
// create a new room and return.
func (server *ChatServer) GetRoom(name string) *Room {
	server.lock.Lock()
	defer server.lock.Unlock()
	if _, ok := server.Rooms[name]; !ok {
		//房号不存在-则创建
		room := &Room{
			Server:  server,             //当前访问
			Name:    name,               //房间名
			lock:    new(sync.RWMutex),  //读写信号
			Clients: make(map[string]*Client),    //创建客户端
			In:      make(chan *Message),        //消息chan
		}
		//启动监听
		go room.Listen()
		server.Rooms[name] = room
	}
	return server.Rooms[name]
}

// This method maybe should add a lock.
//状态监测
func (server *ChatServer) reportStatus() {
	for {
		time.Sleep(time.Second)
		server.lock.RLock()
		for _, room := range server.Rooms {
			fmt.Printf("Status: %s:%d\n", room.Name, len(room.Clients))
		}
		server.lock.RUnlock()
	}
}

//监听访问
func (server *ChatServer) ListenAndServe() {
	listener, err := net.Listen("tcp", server.Bind_to)

	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//服务状态监听
	go server.reportStatus()
	// Main loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %s\n", err.Error())
			time.Sleep(time.Second)
			continue
		}

		//穿件客户端链接
		c := &Client{
			Server: server,
			Name:   conn.RemoteAddr().String(),
			Conn:   conn,
			lock:   new(sync.RWMutex),             //读写信号
			Rooms:  make(map[string]*Room),
			In:     make(chan *Message, 100),
			Out:    make(chan *Message, 100),
			Quit:   make(chan struct{}),
		}
		//监听客户端
		go c.Listen()
		//响应客户端
		go c.Resp()
		//接受访问端
		go c.Recv()
	}
}
