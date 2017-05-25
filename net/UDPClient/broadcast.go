package UDPClient

import (
	"net"
	. "github.com/lflxp/arm/utils"
	"fmt"
	"time"
)

type UDPClient interface {
	Scan()
	SetRemteAddr()
	SetLocalAddr()
}

type Broadcast struct {
/*
	# 局域网广播
	# 实现：如果要让网络（同一网络）中的所有计算机都能收到这个数据包，就应该将这个数据包的接收者地址设置为这个网络中的最高的主机号。通常255.255.255.255就可以达到这个要求。
	# 参考：http://studygolang.com/articles/5233
*/
	Raddr 		net.UDPAddr
	Laddr 		net.UDPAddr
	Net 		string
	Port 		int
}

func (this *Broadcast) Scan() {
	this.SetLocalAddr()
	this.SetRemteAddr()

	conn,err := net.DialUDP(this.Net,&this.Laddr,&this.Raddr)
	CheckError(err)
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

//配置广播包地址255.255.255.255
func (this *Broadcast) SetRemteAddr() {
	this.Raddr = net.UDPAddr{
		IP:net.IPv4(255,255,255,255),
		Port:this.Port,
	}
	//raddr,err := net.ResolveUDPAddr("udp","255.255.255.255:"+strconv.Itoa(this.Port))
	//CheckError(err)
	//this.Raddr = raddr
}
//配置发送端地址
func (this *Broadcast) SetLocalAddr() {
	var local Local
	ip,err := local.GiveOneIp()
	CheckError(err)
	laddr,err := net.ResolveUDPAddr(this.Net,fmt.Sprintf("%s:%d",ip,RandInt(10000,20000)))
	CheckError(err)
	this.Laddr = *laddr
}