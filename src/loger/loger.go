package loger

import (
	"os"
	"fmt"
)

const(
	PORT_NOT_IN_RANGE = "端口超出范围（1-65535）"
	READ_UDP_ERROR = "获取UDP数据出错"
	UUID_USED = "UUID已被占用"
	UUID_UNREGISTED = "UUID未注册"
)


func ErrorString(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func LogErrorString(err string) {
	fmt.Println(err)
}
func LogError(err error) {
	fmt.Println(err)
}
func Info(info string){
	fmt.Println(info)
}

func Used(args ...interface{}) {
	
}