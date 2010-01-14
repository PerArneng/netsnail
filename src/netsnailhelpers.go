package netsnail

import (
	"os"
	"fmt"
	"log"
	"syscall"
	"time"
)

var logger *log.Logger = log.New(os.Stdout, nil, "", log.Lok)

func AbortIfError(err os.Error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "netsnail: aborted: %s\n", err)
		os.Exit(1)
	}
}

func Logf(format string, v ...) {
	now := time.LocalTime().Format("2006-01-02 15:04:05")
	go logger.Logf(now+": "+format, v)
}

func Sleep(ms int) {
	err := syscall.Sleep(int64(ms * 1000000))
	if err != 0 {
		Logf("sleep error: %d\n", err)
	}
}
