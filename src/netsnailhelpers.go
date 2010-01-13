package netsnail

import (
	"os"
	"fmt"
	"log"
)

var logger *log.Logger = log.New(os.Stdout, nil, "netsnail: ", log.Lok)

func AbortIfError(err os.Error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "netsnail: aborted: %s\n", err)
		os.Exit(1)
	}
}

func Logf(format string, v ...) { go logger.Logf(format, v) }
