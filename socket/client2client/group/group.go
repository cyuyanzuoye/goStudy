package client2client

import "net"

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
			panic(err.Error())
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
		return 0, nil
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
