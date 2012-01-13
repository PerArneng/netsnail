/*
 * (C) 2010 Per Arneng
 * License: GPL v2
 */
package netsnail

import (
	"os"
	"fmt"
	"log"
	"syscall"
	"time"
	"net"
)

const NR_NANOSEC_IN_MS int64 = 1000000

var logger *log.Logger = log.New(os.Stdout, "", log.LstdFlags)

func AbortIfError(err os.Error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "netsnail: aborted: %s\n", err)
		os.Exit(1)
	}
}

func Logf(format string, v ...interface{}) {
	now := time.LocalTime().Format("2006-01-02 15:04:05")
	go logger.Printf(now+": "+format, v...)
}

func Sleep(ms int) {
	err := syscall.Sleep(int64(ms) * NR_NANOSEC_IN_MS)
	if err != 0 {
		Logf("sleep error: %d\n", err)
	}
}

func TCPConnect(hostname string, port int) (*net.TCPConn, os.Error) {
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
