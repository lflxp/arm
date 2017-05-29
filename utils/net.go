package utils

import (
	"net"
	"fmt"
	"strings"
	"errors"
)

type Local int

func (this *Local) GetMac() {
	interfaces,err := net.Interfaces()
	CheckError(err)
	for _,inter :=range interfaces {
		mac := inter.HardwareAddr
		name := inter.Name
		fmt.Printf("MAC=%s NAME=%s \n",mac,name)
	}
}

func (this *Local) GetIps() []string {
	result := []string{}
	interfaces,err := net.Interfaces()
	CheckError(err)
	for _,inter :=range interfaces {
		ip,err := inter.Addrs()
		CheckError(err)
		for _,x := range ip {
			//fmt.Printf("Network=%s IP/Net=%s\n",x.Network(),x.String())
			result = append(result,x.String())
		}
	}
	return result
}

func (this *Local) GiveOneIp() (string,error) {
	rs := this.GetIps()
	for _,data := range rs {
		println(data)
		if strings.Contains(data,"24") {
			return strings.Split(data,"/")[0],nil
		}
	}
	return "",errors.New("没有255.255.255.0网段的IP设置")
}