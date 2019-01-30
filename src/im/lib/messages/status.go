package messages

import (
	"bytes"
	"encoding/binary"
)

type AuthenticationStatus struct {
	status int32
	ip int32 //兼容版本0
}

func (auth *AuthenticationStatus) ToData(version int) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, auth.status)
	if version == 0 {
		binary.Write(buffer, binary.BigEndian, auth.ip)
	}
	buf := buffer.Bytes()
	return buf
}

func (auth *AuthenticationStatus) FromData(version int, buff []byte) bool {
	if len(buff) < 4 {
		return false
	}
	buffer := bytes.NewBuffer(buff)
	binary.Read(buffer, binary.BigEndian, &auth.status)
	if version == 0 {
		if len(buff) < 8 {
			return false
		}
		binary.Read(buffer, binary.BigEndian, &auth.ip)
	}
	return true
}

