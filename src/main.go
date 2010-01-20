/*
 * (C) 2010 Per Arneng
 * License: GPL v2
 */
package main

import (
	"fmt"
	"netsnail"
	"net"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	connectionCounter := 0

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

		connectionCounter += 1

		go func() {

			id := fmt.Sprintf("%s[%d]", con.LocalAddr(), connectionCounter)

			netsnail.Logf("%s: connected\n", id)

			proxy, err := netsnail.NewProxy(id, con, conf.Hostname, conf.Port,
				conf.TransferDelay, conf.InitialDelay)
			if err != nil {
				netsnail.Logf("%s: creating proxy failed: %s\n", id, err)
				con.Close()
			} else {
				proxy.Start()
			}

		}()
	}

	listener.Close()

	fmt.Printf("netsnail %s\n", conf)
}
