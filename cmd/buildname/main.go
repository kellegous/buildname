package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kellegous/buildname"
)

func printUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s version\n",
		filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		printUsage()
		os.Exit(1)
	}

	fmt.Println(buildname.FromVersion(flag.Arg(0)))
}
