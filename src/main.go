package main

import (
	"fmt"
	"netsnail"
	"net"
	//"os"
)

func main() {

	conf := netsnail.NewConfig()
	conf.ParseArgs()

	tcpAddress, err :=
		net.ResolveTCPAddr(fmt.Sprintf("localhost:%d", conf.LocalPort))
	netsnail.AbortIfError(err)

	listener, err := net.ListenTCP("tcp", tcpAddress)
	netsnail.AbortIfError(err)

	for {
		con, err := listener.AcceptTCP()
		netsnail.AbortIfError(err)

		proxy, err := netsnail.NewProxy(con, conf.Hostname, conf.Port)
		if err != nil {
			netsnail.Logf("could not create proxy: %s\n", err)
			con.Close()
		} else {
			proxy.Start()
		}
	}

	listener.Close()

	fmt.Printf("netsnail %s\n", conf)
}
