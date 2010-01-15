/*
 * (C) 2010 Per Arneng
 * License: GPL v2
 */
package netsnail

import (
	"fmt"
	"net"
	"os"
)

type DelayProxy struct {
	id            string
	clientConn    *net.TCPConn
	serverConn    *net.TCPConn
	transferDelay int
	initialDelay  int
}

func NewProxy(id string, clientConn *net.TCPConn, hostname string, port int, transferDelay int, initialDelay int) (*DelayProxy, os.Error) {

	addr, err := net.ResolveTCPAddr(fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		return nil, err
	}

	serverConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &DelayProxy{id, clientConn, serverConn, transferDelay, initialDelay}, nil
}

func (this *DelayProxy) Start() { go this.startDataTransfer() }

func (this *DelayProxy) startDataTransfer() {

	Sleep(this.initialDelay)

	finishedChan := make(chan int)

	go this.tcpForward(this.clientConn, this.serverConn, finishedChan)
	go this.tcpForward(this.serverConn, this.clientConn, finishedChan)

	<-finishedChan
	<-finishedChan

	this.Close()

	Logf("%s: connection finished\n", this.id)
}

func (this *DelayProxy) tcpForward(src *net.TCPConn, dest *net.TCPConn, finishedChan chan int) {
	var err os.Error
	var n int
	var buffer = new([500]byte)

	for {
		n, err = src.Read(buffer)
		if n < 1 {
			break
		}

		if err != nil {
			Logf("%s: read error: %d %s\n", this.id, n, err)
			break
		}

		Sleep(this.transferDelay)

		n, err = dest.Write(buffer[0:n])
		if n < 1 {
			break
		}

		if err != nil {
			Logf("%s: write error: %d %s\n", this.id, n, err)
			break
		}
	}

	finishedChan <- 1
}

func (this *DelayProxy) Close() {

	if this.clientConn != nil {
		this.clientConn.Close()
	}

	if this.serverConn != nil {
		this.serverConn.Close()
	}
}
