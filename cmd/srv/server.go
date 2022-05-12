package main

import (
	"goradio/mp3"
	"log"
	"os"
)

func main() {
	var url = "https://stream.klassikradio.de/meditation/mp3-192/radiode/"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}
	player, err := mp3.New(url)
	if err != nil {
		log.Fatal(err)
	}
	defer player.Close()
	player.Play()
	for {
	}
}
