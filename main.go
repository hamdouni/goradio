package main

import (
	"flag"
	"log"
)

const namedpipe = "/tmp/goradio.np"

func main() {
	daemon := flag.Bool("d", false, "daemon server mode")
	flag.Parse()

	if *daemon {
		// mode server : on lit le namedpipe
		if err := server(); err != nil {
			log.Fatalf("server returns %v", err)
		}
	} else {
		// mode cli : on cherche un server ou on en crée un
		if err := client(); err != nil {
			log.Fatalf("client returns %v", err)
		}
	}
}
