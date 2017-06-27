package smode

import (
	"net"
	"packet"
	"loger"
	"time"
)

const(
	STATUS_UNREGISTED = 0x00
	STATUS_REGISTING = 0x01
	STATUS_REGISTED = 0x02
)

var status byte
var conn *net.UDPConn
var respTime time.Time

func Main(uuid string, _conn *net.UDPConn) {

	status = STATUS_UNREGISTED
	conn = _conn

	go regist(uuid)

	// loop to get data
	for {
		var buf [65507]byte
		n, client_addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			loger.LogError(err)
			continue
		}
		if n <=0 {
			loger.LogErrorString(loger.READ_UDP_ERROR)
			continue
		}
		go handlePackage(packet.NewPacket(client_addr, buf[0:n]))
	}
}

func regist(uuid string) {
	for {
		for status = STATUS_REGISTING;status != STATUS_REGISTED; {
			//regist to server
			loger.Info("Registing")
			registPacket := packet.NewRegistPacket(uuid)
			SendPacket(registPacket)
			time.Sleep(3 * time.Second)
		}
		loger.Info("Registed")
		//heartbeat
		ticker := time.NewTicker( time.Second * 10 )
		heartbeatPacket := packet.NewHeartBeatPacket(uuid)
		for status == STATUS_REGISTED {
			<-ticker.C
			SendPacket(heartbeatPacket)
			if time.Now().Sub(respTime).Seconds() > 60 {
				status = STATUS_UNREGISTED
				loger.Info("Server Timeout")
			}
		}
	}
}

func handlePackage(p *packet.Packet) {
	switch p.Type() {
		case packet.OK_PACKET:
			handleOkPacket(p)
		case packet.ERROR_PACKET:
			handleErrorPacket(p)
	}
}

func SendPacket(p packet.BytesPacket) (int, error) {
	n, err := conn.Write(p.Bytes());
	if err != nil {
		loger.LogError(err)
	}
	return n, err
}

func SendPacketToUDP(p packet.BytesPacket, addr *net.UDPAddr) (int, error) {
	n, err := conn.WriteToUDP(p.Bytes(), addr);
	if err != nil {
		loger.LogError(err)
	}
	return n, err
}