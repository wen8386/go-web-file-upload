package common

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"net"
	"os/exec"
)

func ShowPic(picpath string) {
	exec.Command("open", picpath).Start()

}

func GetLocalIpAddress() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	}

	for _, ip := range addrs {
		if ipnet, ok := ip.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil {
				//fmt.Println(ipnet.IP)
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func CreateQR(URL string) {
	qr, err := qrcode.New(URL, qrcode.Medium)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(qr.ToSmallString(false))
	fmt.Println(URL)
}
