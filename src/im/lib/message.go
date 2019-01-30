package lib
/**
  主要封装的是通信的消息机制
 */

//消息构造器
type MessageCreator func()IMessage
var message_creators map[int]MessageCreator = make(map[int]MessageCreator)

type VersionMessageCreator func()IVersionMessage
var vmessage_creators map[int]VersionMessageCreator = make(map[int]VersionMessageCreator)

//消息类型
const  MSG_ACK  = 1




//消息接口
type IMessage interface {
	ToData() []byte
	FromData(buff []byte) bool
}

//消息接接口
type IVersionMessage interface {
	ToData(version int) []byte
	FromData(version int, buff []byte) bool
}


//通用消息处理
type Message struct {
	cmd  int
	seq  int
	version int
	flag int
	body interface{}
}
func (message *Message) ToData() []byte {
	if message.body != nil {
		if m, ok := message.body.(IMessage); ok {
			return m.ToData()
		}
		if m, ok := message.body.(IVersionMessage); ok {
			return m.ToData(message.version)
		}
		return nil
	} else {
		return nil
	}
}

func (message *Message) FromData(buff []byte) bool {
	cmd := message.cmd
	//消息工程构造
	if creator, ok := message_creators[cmd]; ok {
		c := creator()
		r := c.FromData(buff)
		message.body = c
		return r
	}
	if creator, ok := vmessage_creators[cmd]; ok {
		c := creator()
		r := c.FromData(message.version, buff)
		message.body = c
		return r
	}

	return len(buff) == 0
}






