package main

import (
	. "github.com/lflxp/arm/udp/UDPClient"
	. "github.com/lflxp/arm/udp/UDPServer"
	"flag"
)

var stype = flag.String("type","client","Client or Server")

func main() {
	flag.Parse()
	if *stype == "client" {
		DhcpClient()
	} else if *stype == "server" {
		DhcpServer()
	} else {
		println("USAGE:client|server")
	}
}
