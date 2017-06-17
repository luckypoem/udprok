package server

import (
	"flag"
	"loger"
	"net"
	"strconv"
	"time"
	// "fmt"
)

type SModeClient struct {
	uuid string
	addr *net.UDPAddr
	status byte
	updatetime time.Time
}

func clean() {

}

var conn *net.UDPConn
var smodeclients map[string]*SModeClient = make(map[string]*SModeClient)

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
		go handlePackage(&Packet{client_addr,n,buf[0:n]})
	}

}