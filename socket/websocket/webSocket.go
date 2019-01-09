package main

import (
	"fmt"

	"golang.org/x/net/websocket"

	"log"

	"net/http"
)

func Echo(ws *websocket.Conn) {

	var err error
	for {

		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}
		fmt.Println("reveived from client: " + reply)
		msg := "received:" + reply
		fmt.Println("send to client:" + msg)
		//这里是发送消息
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("send failed:", err)
			break
		}
	}
}

func main() {

	http.Handle("/websocket", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
