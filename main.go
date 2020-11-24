package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()

	if len(urls) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if wordlist == "" {
		flag.Usage()
		os.Exit(1)
	}
}
