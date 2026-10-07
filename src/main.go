package main

import (
	"flag"
	_import "lib/src/import"
	"lib/src/tui"
	"lib/src/web"
)

func main() {
	var importFlag, webFlag bool
	flag.BoolVar(&importFlag, "import", false, "Import albums")
	flag.BoolVar(&importFlag, "i", false, "Import albums")
	flag.BoolVar(&webFlag, "web", false, "Start web server")
	flag.BoolVar(&webFlag, "w", false, "Start web server")
	flag.Parse()

	if importFlag {
		_import.Main()
	} else if webFlag {
		web.Main()
	} else {
		tui.TUIMAIN()
	}
}
