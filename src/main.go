package main

import (
	"flag"
	_import "lib/src/import"
	"lib/src/tui"
	"lib/src/web"
)

func main() {
	var importFlag, webFlag, terminalFlag bool
	flag.BoolVar(&importFlag, "import", false, "Import albums")
	flag.BoolVar(&importFlag, "i", false, "Import albums")
	flag.BoolVar(&webFlag, "web", false, "Start web server")
	flag.BoolVar(&webFlag, "w", false, "Start web server")
	flag.BoolVar(&terminalFlag, "terminal", false, "Start terminal ui")
	flag.BoolVar(&terminalFlag, "t", false, "Start terminal ui")
	flag.Parse()

	if importFlag {
		_import.Main()
	} else if webFlag {
		web.Main()
	} else if terminalFlag {
		tui.TUIMAIN()
	} else {
		web.Main()
	}
}
