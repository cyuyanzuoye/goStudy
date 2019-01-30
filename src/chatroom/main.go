// Package main  provides ...
package main

import (
	"runtime"
	"sync"
	"chatroom/chat"
	"fmt"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("chat-room 初始化")
}

func main() {
	server := chat.NewChatServer("127.0.0.1:12345", make(map[string]*chat.Room), new(sync.RWMutex))
	server.ListenAndServe()
}

