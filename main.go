package main

import (
	"github.com/reiver/spacemonkey-init/arg"
	"github.com/reiver/spacemonkey-init/srv/log"

	"fmt"
	"os"
)

func main() {

	var command string = ""
	{
		if 1 <= len(arg.Values) {
			command = arg.Values[0]
		}
	}
	logsrv.Logf("command: %q", command)

	{
		var err error

		switch command {
		default:
			err = do()
		}

		if nil != err {
			fmt.Fprintln(os.Stderr, "error: ", err)
			os.Exit(1)
			return
		}
	}

	fmt.Println("♜ shâh mât")
}
