package main

import (
	"fmt"
	"goradio/cli"
	"goradio/player"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func client(player player.Player) error {

	p := tea.NewProgram(cli.InitModel(player), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	return nil
}
