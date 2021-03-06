package main

import (
	"flag"
	"github.com/automoto/devlog/pkg"
)

func main() {
	outputFlagPtr := flag.String("p", "", "Path to the directory where devlog will save notes to. Only pass in a path to a directory, the file names will be auto-generated.")
	templateFlagPtr := flag.String("t", "", "File path to the .gohtml file for customizing your devlog notes.")
	flag.Parse()
	pkg.Start(*templateFlagPtr, *outputFlagPtr)
}
