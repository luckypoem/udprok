package smode

import (
	"packet"
	"time"
)

func handleOkPacket(p *packet.Packet) {
	switch status {
		case STATUS_REGISTING:
			respTime = time.Now()
			status = STATUS_REGISTED
		case STATUS_REGISTED:
			respTime = time.Now()
	}
}