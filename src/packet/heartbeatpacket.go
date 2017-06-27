package packet

import (
	"bytes"
)

type HeartBeatPacket struct {
	UUID string
}

func NewHeartBeatPacket(uuid string) *HeartBeatPacket {
	return &HeartBeatPacket{uuid}
}

func (packet *HeartBeatPacket)Bytes() []byte {
	var b bytes.Buffer
	b.WriteByte(HEARTBEAT_PACKET)
	b.WriteString(packet.UUID)
	return b.Bytes()
}

func (packet *Packet)HeartBeatPacket() *HeartBeatPacket {
	return &HeartBeatPacket{string(packet.buff[1:])}
}