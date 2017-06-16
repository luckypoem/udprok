package loger

import (
	"os"
	"fmt"
)

const(
	PORT_NOT_IN_RANGE = "端口超出范围（1-65535）"
	READ_UDP_ERROR = "获取UDP数据出错"
)


func Error(err string) {
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

func Used(args ...interface{}) {
	
}