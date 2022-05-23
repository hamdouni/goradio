package main

import (
	"flag"
	"os"
)

func main() {
	var daemonize = flag.Bool("d", false, "daemon")
	flag.Parse()

	print(os.Args[0])
	print(*daemonize)
}
