package messages
//保存在磁盘中但不再需要处理的消息
type IgnoreMessage struct {

}

func (ignore *IgnoreMessage) ToData() []byte {
	return nil
}

func (ignore *IgnoreMessage) FromData(buff []byte) bool {
	return true
}
