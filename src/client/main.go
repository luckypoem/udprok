package client

import (
	"net"
	"loger"
	"fmt"
	"flag"
)

func Main() {

	mode := flag.String("m", "c", "Mode to launch")

	server_addr, _ := net.ResolveUDPAddr("udp", "qcloud.zjxu.org:7450")
	conn, _ := net.DialUDP("udp", nil, server_addr)
	conn.Write([]byte("Hello World!"))

	var buf [20]byte
	n, client_addr, _ := conn.ReadFromUDP(buf[0:])
	fmt.Println("msg is ", string(buf[0:n]))
	loger.Used(n, client_addr, mode)

	if mode == "c" {
		
	}
	// fmt.Println("ip:", client_addr.IP)
	// fmt.Println("port:", client_addr.Port)
}