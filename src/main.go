package main

import (
	"fmt"
	"netsnail"
    "net"
    "os"
)

func main() {

	conf := netsnail.NewConfig()
	conf.ParseArgs()

    tcpAddress, err := 
        net.ResolveTCPAddr(fmt.Sprintf("localhost:%d", conf.LocalPort))
    if err != nil {
        fmt.Printf("failed %s\n", err)
        os.Exit(1)
    }

    listener, err := net.ListenTCP("tcp", tcpAddress)
    if err != nil {
        fmt.Printf("failed %s\n", err)
        os.Exit(1)
    }


    con, err := listener.AcceptTCP()
    
    err := con.Close()


    listener.Close()

	fmt.Printf("netsnail %s\n", conf)
}
