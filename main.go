package main

import "flag"

func main() {
	modePtr := flag.String("mode", "api", "Mode Run (api/cli)")
	flag.Parse()
	if *modePtr == "api" {
		Server().Run()
	} else {
		Console().Run()
	}

}
