package server

import (
	"log"

	"goradio/network/mp3"

	"goradio/shoutcast"
)

type MP3player struct {
	Playing       bool
	URL           string
	Title         string
	Err           error
	NetWorkPlayer *mp3.MP3player
}

func NewMP3player(url string) (player *MP3player, err error) {
	player = new(MP3player)
	player.Playing = false
	player.URL = url
	netmp3, err := mp3.New(url, func(m *shoutcast.Metadata) {
		// callback
		player.Title = m.StreamTitle
	})
	if err != nil {
		return player, err
	}
	player.NetWorkPlayer = netmp3
	return player, nil
}

func (m *MP3player) Pause() {
	m.NetWorkPlayer.Pause()
}

func (m *MP3player) Close() {
	m.Playing = false
	m.NetWorkPlayer.Close()
}

func (m *MP3player) Play() {
	m.Playing = true
	if err := m.NetWorkPlayer.Play(); err != nil {
		log.Printf("Play: %s", err)
	}
}
