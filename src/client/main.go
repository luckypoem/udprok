package client

import (
	"net"
	"loger"
	// "fmt"
	"flag"
	"strconv"
	"github.com/bitly/go-simplejson"
    "io/ioutil"
    "packet"
)

var conn *net.UDPConn
var uuid string

func regist() error {
	packet := packet.NewRegistPacket(uuid)
	conn.Write(packet.Bytes())
	return nil
}
func recieve() {
	for {
		var buff [1024]byte
		n, _, err := conn.ReadFromUDP(buff[0:])
		if err != nil || n <= 0 {
			continue
		}

	}
}

func Main() {

	conf_path := flag.String("c", "udprok.json", "Config file path")
	flag.Parse()

	conf_string, err := ioutil.ReadFile(*conf_path)
	loger.CheckError(err)

	conf_json, err := simplejson.NewJson(conf_string)
	loger.CheckError(err)

	mode := conf_json.Get("mode").MustString()
	host := conf_json.Get("host").MustString()
	port := conf_json.Get("port").MustInt()
	uuid = conf_json.Get("uuid").MustString()

	if port < 1 || port > 65535 {
		loger.ErrorString(loger.PORT_NOT_IN_RANGE)
	}

	server_addr, err := net.ResolveUDPAddr("udp", host + ":" + strconv.Itoa(port))
	loger.CheckError(err)

	_conn, err := net.DialUDP("udp", nil, server_addr)
	loger.CheckError(err)
	conn = _conn

	if mode == "server" {
		regist()
	}
}