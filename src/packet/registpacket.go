package packet

import (
	"bytes"
)

type RegistPacket struct {
	UUID string
}

func NewRegistPacket(uuid string) *RegistPacket {
	return &RegistPacket{uuid}
}

func (packet *RegistPacket)Bytes() []byte {
	var b bytes.Buffer
	b.WriteByte(REGIST_PACKET)
	b.WriteString(packet.UUID)
	return b.Bytes()
}

func (packet *Packet)RegistPacket() *RegistPacket {
	return &RegistPacket{string(packet.buff[1:])}
}