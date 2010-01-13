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
	flag.IntVar(&this.LocalPort, "local-port", 9091, "the port of this proxy")
	flag.IntVar(&this.Port, "port", 80, "the port of the remote host")
	flag.StringVar(&this.Hostname,
		"hostname", "localhost", "the remote hostname")
	flag.IntVar(&this.TransferDelay, "transfer-delay", 0, "the delay on data transfer in ms")
	flag.Parse()
}

func (this *Config) String() string {
	return fmt.Sprintf("[Config| local-port: %d, hostname: '%s', port: %d, transfer-delay: %d]",
		this.LocalPort, this.Hostname, this.Port, this.TransferDelay)
}
