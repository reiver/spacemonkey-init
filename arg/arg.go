package arg

import (
	"flag"
)

var (
	Values []string
)

var (
	Verbose bool
)

func init() {
	flag.BoolVar(&Verbose, "v", false, "verbose logs outputted")

	flag.Parse()

	Values = flag.Args()
}
