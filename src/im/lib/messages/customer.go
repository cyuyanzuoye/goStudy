package messages

import (
	"bytes"
	"encoding/binary"
)


type CustomerMessage struct {
	customer_appid int64//顾客id所在appid
	customer_id    int64//顾客id
	store_id	   int64
	seller_id	   int64
	timestamp	   int32
	content		   string
}

func (cs *CustomerMessage) ToData() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, cs.customer_appid)
	binary.Write(buffer, binary.BigEndian, cs.customer_id)
	binary.Write(buffer, binary.BigEndian, cs.store_id)
	binary.Write(buffer, binary.BigEndian, cs.seller_id)
	binary.Write(buffer, binary.BigEndian, cs.timestamp)
	buffer.Write([]byte(cs.content))
	buf := buffer.Bytes()
	return buf
}

func (cs *CustomerMessage) FromData(buff []byte) bool {
	if len(buff) < 36 {
		return false
	}
	buffer := bytes.NewBuffer(buff)
	binary.Read(buffer, binary.BigEndian, &cs.customer_appid)
	binary.Read(buffer, binary.BigEndian, &cs.customer_id)
	binary.Read(buffer, binary.BigEndian, &cs.store_id)
	binary.Read(buffer, binary.BigEndian, &cs.seller_id)
	binary.Read(buffer, binary.BigEndian, &cs.timestamp)

	cs.content = string(buff[36:])

	return true
}