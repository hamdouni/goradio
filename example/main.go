package main

import (
	"flag"
	"fmt"
	"log"

	"goradio/network/mp3"
	"goradio/shoutcast"
)

var station = flag.String("station", "https://radio.barim.us/stream", "url of a stream radio")

func main() {
	flag.Parse()
	log.Printf("Playing %s\n", *station)
	player, err := mp3.New(*station, func(m *shoutcast.Metadata) {
		println("Now listening to: ", m.StreamTitle)
	})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if err := player.Play(); err != nil {
		log.Fatalf("error: %s", err)
	}
	defer player.Close()

	println("press enter to quit...")
	fmt.Scanln()
}
