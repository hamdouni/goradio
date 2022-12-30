package main

import (
	"flag"
	"fmt"
	"goradio/network/mp3"
	"log"
)

var station = flag.String("station", "https://radio.barim.us/stream", "url of a stream radio")

func main() {
	flag.Parse()
	log.Printf("Playing %s\n", *station)
	player, err := mp3.New(*station)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	player.Play()
	defer player.Close()

	println("press enter to quit...")
	fmt.Scanln()
}
