package packet

import (
	"bytes"
)

type OkPacket struct {
}

func NewOkPacket() *OkPacket {
	return &OkPacket{}
}

func (packet *OkPacket)Bytes() []byte {
	var b bytes.Buffer
	b.WriteByte(OK_PACKET)
	return b.Bytes()
}

func (packet *Packet)OkPacket() *OkPacket {
	return &OkPacket{}
}