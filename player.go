package main

import (
	"fmt"
	"goradio/mp3"

	tea "github.com/charmbracelet/bubbletea"
)

var player *mp3.MP3player

type playerError struct {
	err error
}

func (e playerError) Error() string { return e.err.Error() }

type playerLoaded struct{}

func load(uri string) tea.Cmd {
	return func() tea.Msg {
		if player != nil {
			player.Close()
		}
		p, err := mp3.New(uri)
		if err != nil {
			return playerError{err: err}
		}
		player = p
		return playerLoaded{}
	}
}

type playerStarted struct {
	status string
}

func play(name string) tea.Cmd {
	return func() tea.Msg {
		if player == nil {
			return playerError{err: fmt.Errorf("player not loaded")}
		}
		if err := player.Play(); err != nil {
			return playerError{err: err}
		}
		return playerStarted{status: fmt.Sprintf("playing %s", name)}
	}
}
