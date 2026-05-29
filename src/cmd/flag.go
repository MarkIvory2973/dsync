package cmd

import (
	"flag"
)

func init() {
	flag.Parse()
}

func GetSrcRoot() string {
	srcRoot := flag.Arg(0)

	return srcRoot
}

func GetDstRoot() string {
	dstRoot := flag.Arg(1)

	return dstRoot
}
