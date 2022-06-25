package server

import (
	"goradio/mp3"
	"goradio/player"
	"log"
)

func Run(responder player.Responder) error {

	var mp3player *mp3.MP3player

	defer responder.Close()

	var loop = true
	for loop {
		buf := responder.ReadRequest()

		// ici on doit interpreter le buffer
		switch buf[0] {
		case 'q':
			loop = false
		case 'z':
			if mp3player != nil {
				mp3player.Paused = !mp3player.Paused
			}
		case 's':
			if mp3player != nil {
				var st string
				if mp3player.Playing {
					st = "play"
				} else if mp3player.Paused {
					st = "pause"
				}
				responder.Write(st)
			} else {
				responder.Write("none")
			}
		case 'u':
			if mp3player != nil {
				responder.Write(mp3player.URL)
			} else {
				responder.Write("none")
			}
		case 'e':
			if mp3player != nil && mp3player.Err != nil {
				responder.Write(mp3player.Err.Error())
			} else {
				responder.Write("none")
			}
		case 'p':
			if mp3player != nil && mp3player.Playing {
				mp3player.Close()
			}
			// le paramètre commence en position 1
			url := buf[1:]
			var err error
			if mp3player, err = mp3.New(url); err != nil {
				// @TODO: write err in response
				mp3player.Err = err
				log.Printf("mp3 err: %s", err)
				continue
			}
			mp3player.Err = nil
			mp3player.Play()
		}
	}
	return nil
}
