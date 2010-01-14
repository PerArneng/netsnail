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
}

func NewConfig() *Config { return new(Config) }

func (this *Config) ParseArgs() {
	flag.IntVar(&this.LocalPort, "l", 9091, "the local port of this proxy")
	flag.IntVar(&this.Port, "r", 80, "the port of the remote host")
	flag.StringVar(&this.Hostname,
		"h", "localhost", "the remote hostname")
	flag.IntVar(&this.TransferDelay, "d", 0,
		"the delay on data transfer in ms")
	flag.Parse()
}

func (this *Config) String() string {
	return fmt.Sprintf("[Config| local-port: %d, hostname: '%s', port: %d, transfer-delay: %d]",
		this.LocalPort, this.Hostname, this.Port, this.TransferDelay)
}
