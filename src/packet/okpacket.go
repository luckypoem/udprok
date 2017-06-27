package packet

import (
	"bytes"
)

var _okPkg *OkPacket = nil
type OkPacket struct {
}

func NewOkPacket() *OkPacket {
	if _okPkg == nil {
		_okPkg = &OkPacket{}
	}
	return _okPkg
}

func (packet *OkPacket)Bytes() []byte {
	var b bytes.Buffer
	b.WriteByte(OK_PACKET)
	return b.Bytes()
}

func (packet *Packet)OkPacket() *OkPacket {
	return &OkPacket{}
}