package server

import (
	gonet "goradio/network/mp3"
)

type MP3player struct {
	Playing       bool
	URL           string
	Err           error
	NetWorkPlayer *gonet.MP3player
}

func NewMP3player(url string) (mp3 *MP3player, err error) {
	mp3 = new(MP3player)
	mp3.Playing = false
	mp3.URL = url
	netmp3, err := gonet.New(url)
	if err != nil {
		return mp3, err
	}
	mp3.NetWorkPlayer = netmp3
	return mp3, nil
}

func (m *MP3player) Close() {
	m.Playing = false
	m.NetWorkPlayer.Close()
}

func (m *MP3player) Play() {
	m.Playing = true
	m.NetWorkPlayer.Play()
}
