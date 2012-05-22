/*
 * (C) 2010 Per Arneng
 * License: GPL v2
 */
package netsnail

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const NR_NANOSEC_IN_MS int64 = 1000000

var logger *log.Logger = log.New(os.Stdout, "", log.LstdFlags)

func AbortIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "netsnail: aborted: %s\n", err)
		os.Exit(1)
	}
}

func Logf(format string, v ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	go logger.Printf(now+": "+format, v...)
}

func Sleep(ms int) {
	time.Sleep(time.Duration(int64(ms) * NR_NANOSEC_IN_MS))
}

func TCPConnect(hostname string, port int) (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
