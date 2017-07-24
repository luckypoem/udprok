package packet

import (
	"bytes"
)

type ConnectPacket struct {
	UUID string
}

func NewConnectPacket(uuid string) *ConnectPacket {
	return &ConnectPacket{uuid}
}

func (packet *ConnectPacket)Bytes() []byte {
	var b bytes.Buffer
	b.WriteByte(CONNECT_PACKET)
	b.WriteString(packet.UUID)
	return b.Bytes()
}

func (packet *Packet)ConnectPacket() *ConnectPacket {
	return &ConnectPacket{string(packet.buff[1:])}
}