package udprokc

import (
	"flag"
	"loger"
	"net"
	"strconv"
	"packet"
	"time"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

const(
	STATUS_UNCONN = 0x00
	STATUS_GETINFO = 0x01
	STATUS_INFOED = 0x02
	STATUS_CONNING = 0x03
	STATUS_CONNED = 0x04
)

var status byte
var conn *net.UDPConn

func Main(){

	configPath := flag.String("c", "udproks.json", "Config file path")
	flag.Parse()

	configData, err := ioutil.ReadFile(*configPath)
	loger.CheckError(err)

	configJson, err := simplejson.NewJson(configData)
	loger.CheckError(err)

	host := configJson.Get("host").MustString()
	port := configJson.Get("port").MustInt()
	uuid := configJson.Get("uuid").MustString()

	if port < 1 || port > 65535 {
		loger.ErrorString(loger.PORT_NOT_IN_RANGE)
	}

	serverAddr, err := net.ResolveUDPAddr("udp", host + ":" + strconv.Itoa(port))
	loger.CheckError(err)

	_conn, err := net.DialUDP("udp", nil, serverAddr)
	loger.CheckError(err)

	conn = _conn

	status = STATUS_UNCONN

	go connect(uuid)

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

func connect(uuid string) {
	connectPacket := packet.NewConnectPacket(uuid)
	for {
		//get smode info from server
		for status = STATUS_CONNING; status != STATUS_CONNED; {
			loger.Info("connecting")
			SendPacket(connectPacket)
			time.Sleep(3 * time.Second)
		}
		loger.Info("connected")
	}
}