package server

import (
	"flag"
	"loger"
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	
}

func Main() {

	port := flag.Int("p", 7450, "The port to listen UDP")
	addr := flag.String("l", "0.0.0.0", "The address to listen UDP")
	flag.Parse()

	if *port < 1 || *port > 65535 {
		loger.Error(loger.PORT_NOT_IN_RANGE)
	}

	server_addr, err := net.ResolveUDPAddr("udp", *addr + ":" + strconv.Itoa(*port))
	loger.CheckError(err)

	conn, err := net.ListenUDP("udp", server_addr)
	loger.CheckError(err)

	// var buf [20]byte
	// n, client_addr, _ := conn.ReadFromUDP(buf[0:])
	// fmt.Println("msg is ", string(buf[0:n]))
	// fmt.Println("ip:", client_addr.IP)
	// fmt.Println("port:", client_addr.Port)
    // _, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
}