/*
 * (C) 2010 Per Arneng
 * License: GPL v2
 */
package netsnail

import (
	"fmt"
	"flag"
)

type Config struct {
	LocalPort     int
	Port          int
	Hostname      string
	TransferDelay int
	InitialDelay  int
}

func NewConfig() *Config { return new(Config) }

func (this *Config) ParseArgs() {
	flag.IntVar(&this.LocalPort, "l", 9091, "the local port of this proxy")
	flag.IntVar(&this.Port, "r", 80, "the port of the remote host")
	flag.StringVar(&this.Hostname,
		"h", "localhost", "the remote hostname")
	flag.IntVar(&this.TransferDelay, "d", 0,
		"the delay on data transfer in ms")
	flag.IntVar(&this.InitialDelay, "i", 0,
		"a fixed delay in ms when each connection is made")
	flag.Parse()
}

func (this *Config) String() string {
	return fmt.Sprintf("[Config| local-port: %d, hostname: '%s', port: %d, transfer-delay: %d]",
		this.LocalPort, this.Hostname, this.Port, this.TransferDelay)
}
