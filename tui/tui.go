package tui

import (
	"goradio/player"

	tea "github.com/charmbracelet/bubbletea"
)

// Run the terminal user interface
func Run(player player.Player) error {
	p := tea.NewProgram(initModel(player), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}
