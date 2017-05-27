package UDPClient
/*
1.协议标准:
客户端常用端口：19878
UDP发送广播包端口：10000到20000之间
*/
import (
	"net"
	. "github.com/lflxp/arm/utils"
	"fmt"
	"time"
	"github.com/lflxp/arm/udp"
)

type UDPClient interface {
	//发送广播包 扫描服务器
	Discover()
	//获得连接
	GetConnect()
	//关闭
	Close()
	GetOffer()
}

type Broadcast struct {
/*
	# 局域网广播
	# 实现：如果要让网络（同一网络）中的所有计算机都能收到这个数据包，就应该将这个数据包的接收者地址设置为这个网络中的最高的主机号。通常255.255.255.255就可以达到这个要求。
	# 参考：http://studygolang.com/articles/5233
*/
	Conn		*net.UDPConn
	Raddr 		net.UDPAddr
	Laddr 		net.UDPAddr
	//网络模式 udp udp4 udp6
	Net 		string
	//服务器 UDP端口
	Port 		int
}

func (this *Broadcast) GetConnect() {
	this.SetLocalAddr()
	this.SetRemteAddr()
	//发送连接请求
	conn,err := net.DialUDP(this.Net,&this.Laddr,&this.Raddr)
	CheckError(err)
	this.Conn = conn
}
/*
Close  Connections
 */
func (this *Broadcast) Close() error {
	if this.Conn != nil {
		return this.Conn.Close()
	}
	return nil
}

func (this *Broadcast) Discover(packet udp.Packet) {


	daytime := time.Now().String()
	this.Conn.Write([]byte(daytime))
}

//配置广播包地址255.255.255.255
func (this *Broadcast) SetRemteAddr() {
	this.Raddr = net.UDPAddr{
		IP:net.IPv4(255,255,255,255),
		Port:this.Port,
	}
	//raddr,err := udp.ResolveUDPAddr("udp","255.255.255.255:"+strconv.Itoa(this.Port))
	//CheckError(err)
	//this.Raddr = raddr
}
//配置发送端地址
func (this *Broadcast) SetLocalAddr() {
	var local Local
	ip,err := local.GiveOneIp()
	CheckError(err)
	laddr,err := net.ResolveUDPAddr(this.Net,fmt.Sprintf("%s:%d",ip,RandInt(10000,19877)))
	CheckError(err)
	this.Laddr = *laddr
}