package main

import (
	"flag"
	_import "lib/src/import"
	"lib/src/tui"
)

func main() {
	var importFlag bool
	flag.BoolVar(&importFlag, "import", false, "Import albums")
	flag.BoolVar(&importFlag, "i", false, "Import albums")
	flag.Parse()

	if importFlag {
		_import.Main()
	} else {
		tui.TUIMAIN()
	}
}
