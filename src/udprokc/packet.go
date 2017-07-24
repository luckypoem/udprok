package udprokc

import(
	"packet"
	"net"
	"loger"
)

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

func handlePackage(p *packet.Packet) {
	switch p.Type() {
	}
}
