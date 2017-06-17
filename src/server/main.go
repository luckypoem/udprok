package server

import (
	"flag"
	"loger"
	"net"
	"strconv"
	"time"
	"fmt"
)

type SModeClient struct {
	uuid string
	addr *net.UDPAddr
	status byte
	updatetime time.Time
}

type Packet struct {
	client_addr *net.UDPAddr
	n int
	buff []byte
}

const(
	REGIST_PACKET = 0x01
	HEARTBEAT_PACKET = 0x02
	CONNECT_PACKET = 0x03
)

func handlePackage(packet *Packet) {
	switch packet.buff[0] {
		case REGIST_PACKET:
			packet.regist()
		case HEARTBEAT_PACKET:
			packet.heartbeat()
		case CONNECT_PACKET:
			packet.connect()
	}
}

func (packet *Packet)regist() {
	uuid := string(packet.buff[1:])
	smodeclients[uuid] = SModeClient{
		uuid,
		packet.client_addr,
		0,
		time.Now()
	}
	fmt.Println("regist",uuid)
}

func (packet *Packet)heartbeat() {
	uuid := string(packet.buff[1:])
	fmt.Println("heartbeat",uuid)
}

func (packet *Packet)connect() {
	uuid := string(packet.buff[1:])
	fmt.Println("connect",uuid)
}

func clean() {

}

var conn *net.UDPConn
var smodeclients map[*SModeClient]string

func Main() {

	port := flag.Int("p", 7450, "The port to listen UDP")
	addr := flag.String("l", "0.0.0.0", "The address to listen UDP")
	flag.Parse()

	if *port < 1 || *port > 65535 {
		loger.ErrorString(loger.PORT_NOT_IN_RANGE)
	}

	server_addr, err := net.ResolveUDPAddr("udp", *addr + ":" + strconv.Itoa(*port))
	loger.CheckError(err)

	smodeclients = make(map[*SModeClient]string)
	_conn, err := net.ListenUDP("udp", server_addr)
	loger.CheckError(err)
	conn = _conn

	go clean()
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
		go handlePackage(&Packet{client_addr,n,buf[0:n]})
	}

}