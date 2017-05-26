package UDPServer

import (
	"net"
	"time"
	. "github.com/lflxp/arm/utils"
	"fmt"
)

func Server() {
	service := ":1200"
	udpAddr,err := net.ResolveUDPAddr("udp",service)
	CheckError(err)
	conn,err := net.ListenUDP("udp",udpAddr)
	CheckError(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_,addr,err := conn.ReadFromUDP(buf[0:])
	CheckError(err)
	fmt.Printf("read msg:%v\n",string(buf[0:]))
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime),addr)
}

