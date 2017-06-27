package packet

import (
	"bytes"
)

type ErrorPacket struct {
	Msg string
}

func NewErrorPacket(msg string) *ErrorPacket {
	return &ErrorPacket{msg}
}

func (packet *ErrorPacket)Bytes() []byte {
	var b bytes.Buffer
	b.WriteByte(ERROR_PACKET)
	b.WriteString(packet.Msg)
	return b.Bytes()
}

func (packet *Packet)ErrorPacket() *ErrorPacket {
	return &ErrorPacket{string(packet.buff[1:])}
}