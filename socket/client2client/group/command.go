package client2client

import (
	"encoding/json"
	"fmt"
	"io"
	"syscall"
)

//解析
type Parse interface {
	Parse() error
}

//执行命令
type Exec interface {
	Exec() error
}

//退出
type Exit interface {
	Exit() error
}

//发送
type Send interface {
	Send() error
}

//命令内容
type Content struct {
	From    string
	To      string
	Message string
}

//命令
type Command struct {
	Cmd string
	Content
}

//命令代理对象
type Proxy struct {
	cmd interface{}
}

func Parse(data []byte) *Command {
	commandMap := make(map[string]string)
	err := json.Unmarshal(data, commandMap)
	if err != nil {
		fmt.Println("命令解析失败")
	}
	return &Command{Cmd: commandMap["cmd"], Content: commandMap["Content"]}
}

func (cmd *Command) Exit() {

}
