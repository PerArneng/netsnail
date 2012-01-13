/*
 * (C) 2010 Per Arneng
 * License: GPL v2
 */
package netsnail

import "net"

const CHUNK_SIZE int = 1024

type DelayProxy struct {
	id            string
	clientConn    *net.TCPConn
	serverConn    *net.TCPConn
	transferDelay int
	initialDelay  int
}

func NewProxy(id string, clientConn *net.TCPConn, hostname string, port int, transferDelay int, initialDelay int) (*DelayProxy, error) {

	serverConn, err := TCPConnect(hostname, port)
	if err != nil {
		return nil, err
	}

	return &DelayProxy{id, clientConn, serverConn, transferDelay, initialDelay}, nil
}

func (this *DelayProxy) Start() { go this.startProxy() }

func (this *DelayProxy) startProxy() {

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
	var err error
	var n int
	var buffer = make([]byte, CHUNK_SIZE)

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
