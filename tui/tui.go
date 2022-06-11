package tui

import (
	"goradio/player"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jamesnetherton/m3u"
)

// Run the terminal user interface
func Run(player player.Player,playlist m3u.Playlist) error {
	p := tea.NewProgram(initModel(player,playlist), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}
