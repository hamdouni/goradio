package main

import (
	"flag"
	"goradio/com"
	"log"
	"os"
	"os/exec"
)

func main() {
	daemon := flag.Bool("d", false, "daemon server mode")
	flag.Parse()

	pipeplayer, created, err := com.NewPipePlayer()
	if err != nil {
		log.Fatalf("com.NewPipePlayer returns %v", err)
	}

	if *daemon {
		// mode server : on lit le namedpipe
		if err := server(pipeplayer); err != nil {
			log.Fatalf("server returns %v", err)
		}
		os.Exit(0)
	}

	// mode cli
	if created {
		// launch server
		log.Println("client: launching server")
		cmd := exec.Command(os.Args[0], "-d")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Start(); err != nil {
			log.Fatalf("launching server returns %v", err)
		}
	}
	if err := client(pipeplayer); err != nil {
		log.Fatalf("client returns %v", err)
	}
}
