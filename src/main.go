package main

import (
	"fmt"
	"netsnail"
)

func main() {

	conf := netsnail.NewConfig()
	conf.ParseArgs()

	fmt.Printf("netsnail %s\n", conf)
}
