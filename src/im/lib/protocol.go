package lib
/**
 * 端口协议解析
 */
import (
	"io"
	"encoding/hex"
	"encoding/binary"
	"bytes"
	"errors"
	"log"
	"fmt"
)

/**
 * 约定协议头
 */
func WriteHeader(len int32, seq int32, cmd byte, version byte, flag byte, buffer io.Writer) {
	binary.Write(buffer, binary.BigEndian, len)
	binary.Write(buffer, binary.BigEndian, seq)
	t := []byte{cmd, byte(version), flag, 0}
	buffer.Write(t)
}

//读取头部信息
func ReadHeader(buff []byte) (int, int, int, int, int) {
	var length int32
	var seq int32
	buffer := bytes.NewBuffer(buff)
	binary.Read(buffer, binary.BigEndian, &length)
	binary.Read(buffer, binary.BigEndian, &seq)
	cmd, _ := buffer.ReadByte()
	version, _ := buffer.ReadByte()
	flag, _ := buffer.ReadByte()
	return int(length), int(seq), int(cmd), int(version), int(flag)
}

//写入消息
func WriteMessage(w *bytes.Buffer, msg *Message) {
	body := msg.ToData()
	WriteHeader(int32(len(body)), int32(msg.seq), byte(msg.cmd), byte(msg.version), byte(msg.flag), w)
	w.Write(body)
}

//发送消息
func SendMessage(conn io.Writer, msg *Message) error {
	buffer := new(bytes.Buffer)
	WriteMessage(buffer, msg)
	buf := buffer.Bytes()
	n, err := conn.Write(buf)
	if err != nil {
		log.Println("sock write error:", err)
		return err
	}
	if n != len(buf) {
		log.Println("write less:%d %d", n, len(buf))
		return errors.New("write less")
	}
	return nil
}

//有效限制的接受消息
func ReceiveLimitMessage(conn io.Reader, limit_size int, external bool) *Message {
	buff := make([]byte, 12)
	//将数据读取到buff中，一次性处理12个字节
	_, err := io.ReadFull(conn, buff)
	if err != nil {
		log.Println("sock read error:", err)
		return nil
	}
	length, seq, cmd, version, flag := ReadHeader(buff)
	//超出限制
	if length < 0 || length >= limit_size {
		log.Println("invalid len:", length)
		return nil
	}
	fmt.Println(string(buff))
	//没有超出显示
	buff = make([]byte, length)
	_, err = io.ReadFull(conn, buff)
	if err != nil {
		fmt.Println("sock read error:", err)
		return nil
	}

	message := new(Message)
	message.cmd = cmd
	message.seq = seq
	message.version = version
	message.flag = flag
	fmt.Println(string(buff) )
	if !message.FromData(buff) {
		fmt.Println("parse error:%d, %d %d %d %s", cmd, seq, version,
			flag, hex.EncodeToString(buff))
		return nil
	}

	return message
}


func ReceiveMessage(conn io.Reader) *Message {
	return ReceiveLimitMessage(conn, 32*1024, false)
}

//接受客户端消息(external messages)
func ReceiveClientMessage(conn io.Reader) *Message {
	return ReceiveLimitMessage(conn, 32*1024, true)
}

//消息大小限制在1M
func ReceiveStorageSyncMessage(conn io.Reader) *Message {
	return ReceiveLimitMessage(conn, 32*1024*1024, false)
}



