package main

import (
	"fmt"
	"goradio/cli"
	"goradio/player"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func client(player player.Player) error {

	// p := tea.NewProgram(cli.InitModel(player), tea.WithAltScreen())
	p := tea.NewProgram(cli.InitModel(player))
	if err := p.Start(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	// log.Print("client: ask server to quit")
	// pipe.Write([]byte("q"))
	// log.Print("client: ask server to play")
	// pipe.Write([]byte("phttps://stream.klassikradio.de/meditation/mp3-192/radiode/"))
	// log.Print("client: ask server to another stream")
	// pipe.Write([]byte("phttp://37.187.93.104:8097/stream/1/"))
	// time.Sleep(5 * time.Second)
	return nil
}
