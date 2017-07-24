package udproks

import(
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
	STATUS_UNREGISTED = 0x00
	STATUS_REGISTING = 0x01
	STATUS_REGISTED = 0x02
)

var status byte
var conn *net.UDPConn
var respTime time.Time

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

	status = STATUS_UNREGISTED

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

	registPacket := packet.NewRegistPacket(uuid)
	heartbeatPacket := packet.NewHeartBeatPacket(uuid)

	for {
		for status = STATUS_REGISTING;status != STATUS_REGISTED; {
			//regist to server
			loger.Info("Registing")
			SendPacket(registPacket)
			time.Sleep(3 * time.Second)
		}
		loger.Info("Registed")
		//heartbeat
		ticker := time.NewTicker( time.Second * 10 )
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