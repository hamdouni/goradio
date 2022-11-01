package main

import (
	"fmt"
	"goradio/network/mp3"
	"log"
)

func main() {

	player, err := mp3.New("http://51.255.235.165:5068/stream/1/")
	// player, err := mp3.New("http://radio.barim.us/stream")
	if err != nil {
		log.Fatal(err)
	}
	player.Play()
	defer player.Close()

	println("press enter to quit...")
	fmt.Scanln()
}
