package udprokd

import (
	"flag"
	"loger"
	"net"
	"strconv"
	"time"
	"packet"
	"sync"
	"bytes"
)

type Udproks struct {
	UUID string
	Addr *net.UDPAddr
	UpdateTime time.Time
}

var conn *net.UDPConn
var udproksMap map[string]*Udproks = make(map[string]*Udproks)
var mapmutex sync.RWMutex

func Main() {

	// get config
	port := flag.Int("p", 7450, "The port to listen UDP")
	addr := flag.String("l", "0.0.0.0", "The address to listen UDP")
	flag.Parse()

	// check config
	if *port < 1 || *port > 65535 {
		loger.ErrorString(loger.PORT_NOT_IN_RANGE)
	}

	server_addr, err := net.ResolveUDPAddr("udp", *addr + ":" + strconv.Itoa(*port))
	loger.CheckError(err)
	
	_conn, err := net.ListenUDP("udp", server_addr)
	loger.CheckError(err)
	conn = _conn

	// start a thread to timely clean up timeout clients
	go clean()

	// loop to get data
	for {
		var buf [1024]byte
		n, client_addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			loger.LogError(err)
			continue
		}
		if n <=0 {
			loger.LogErrorString(loger.READ_UDP_ERROR)
			continue
		}
		if n >= 1024 {
			n = 1024
		}
		go handlePackage(packet.NewPacket(client_addr, buf[0:n]))
	}

}

func handlePackage(p *packet.Packet) {
	switch p.Type() {
		case packet.REGIST_PACKET:
			registServer(p)
		case packet.HEARTBEAT_PACKET:
			handleHeartbeat(p)
	}
}

func sendPackage(p packet.BytesPacket, addr *net.UDPAddr) (int, error){
	n, err := conn.WriteToUDP(p.Bytes(), addr)
	if err != nil {
		loger.LogError(err)
	}
	return n,err
}

func registServer(p *packet.Packet) {
	uuid := p.RegistPacket().UUID

	mapmutex.Lock()
	client, ok := udproksMap[uuid]

	//check if the uuid has been registed
	if ok {
		if !bytes.Equal(client.Addr.IP, p.Addr.IP) || client.Addr.Port != p.Addr.Port {
			mapmutex.Unlock()
			errorPacket := packet.NewErrorPacket(loger.UUID_USED)
			sendPackage(errorPacket, p.Addr)
		}else{
			client.UpdateTime = time.Now()
			mapmutex.Unlock()
			okPacket := packet.NewOkPacket()
			sendPackage(okPacket, p.Addr)
		}
	}else{
		client = &Udproks{uuid, p.Addr, time.Now()}
		udproksMap[uuid] = client
		mapmutex.Unlock()
		okPacket := packet.NewOkPacket()
		sendPackage(okPacket, p.Addr)
	}
}

func handleHeartbeat(p *packet.Packet) {
	heartbeatPacket := p.HeartBeatPacket()
	uuid := heartbeatPacket.UUID
	mapmutex.RLock()
	client, ok := udproksMap[uuid]
	if ok {
		client.UpdateTime = time.Now()
		mapmutex.RUnlock()
		sendPackage(packet.NewOkPacket(), p.Addr)
	}else{
		mapmutex.RUnlock()
		errorPacket := packet.NewErrorPacket(loger.UUID_UNREGISTED)
		sendPackage(errorPacket, p.Addr)
	}
}

func clean() {
	//clean the registed smode clients which had been timeout.
	ticker := time.NewTicker( time.Second * 10 )
	for {
		<-ticker.C
		now := time.Now()
		mapmutex.Lock()
		for uuid ,client := range udproksMap{
			if now.Sub(client.UpdateTime).Seconds() > 60 {
				delete(udproksMap, uuid)
			}
	    }
	    mapmutex.Unlock()
	}
}