package main

import (
	"goradio/mp3"
	"os"
)

func server() error {

	var player *mp3.MP3player
	var playing bool

	pipe, err := os.OpenFile(namedpipe, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		return err
	}
	defer func() {
		pipe.Close()
		os.Remove(namedpipe)
	}()
	var loop = true
	for loop {
		buf := make([]byte, 255)
		_, err := pipe.Read(buf)
		if err != nil {
			return err
		}

		// ici on doit interpreter le buffer
		switch buf[0] {
		case 'q':
			loop = false
		case 'z':
			if player != nil {
				player.Paused = !player.Paused
			}
		case 'p':
			if playing {
				player.Close()
				playing = false
			}
			url := ""
			// le paramètre va de position 1 jusqu'à un \n
			for i := 1; i < len(buf); i++ {
				if buf[i] == '\n' {
					break
				}
				url += string(buf[i])
			}
			player, err = mp3.New(url)
			if err != nil {
				return err
			}
			defer player.Close()
			playing = true
			player.Play()
		}
	}
	return nil
}
