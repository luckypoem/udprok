package udproks

import (
	"packet"
	"loger"
)

func handleErrorPacket(p *packet.Packet) {
	errorPacket := p.ErrorPacket()
	switch status {
		case STATUS_REGISTING:
			loger.ErrorString(errorPacket.Msg)
		case STATUS_REGISTED:
			loger.LogErrorString(errorPacket.Msg)
			status = STATUS_UNREGISTED
	}
}