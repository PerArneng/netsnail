package netsnail

import (
	"fmt"
	"net"
	"os"
)

type DelayProxy struct {
	clientConn *net.TCPConn
	serverConn *net.TCPConn
}

func NewProxy(clientConn *net.TCPConn, hostname string, port int) (*DelayProxy, os.Error) {

	addr, err := net.ResolveTCPAddr(fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		return nil, err
	}

	serverConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &DelayProxy{clientConn, serverConn}, nil
}

func (this *DelayProxy) Start() {}

func (this *DelayProxy) Close() {

	if this.clientConn != nil {
		this.clientConn.Close()
	}

	if this.serverConn != nil {
		this.serverConn.Close()
	}
}
