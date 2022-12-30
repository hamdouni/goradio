package server

import (
	"goradio/player"
	"log"
)

func Run(responder player.Responder) error {

	var mp3player *MP3player

	defer responder.Close()

	var msg string
	var loop = true
	for loop {
		buf := responder.ReadRequest()

		if len(buf) == 0 {
			continue
		}

		// ici on doit interpreter le buffer
		switch buf[0] {
		case 'q':
			loop = false
		case 'z':
			if mp3player != nil {
				if mp3player.Playing {
					mp3player.Close()
				} else {
					actual_url := mp3player.URL
					var err error
					if mp3player, err = NewMP3player(actual_url); err != nil {
						// @TODO: write err in response
						mp3player.Err = err
						log.Printf("mp3 err: %s", err)
						continue
					}
					mp3player.Err = nil
					mp3player.Play()
				}
			}
		case 's':
			if mp3player != nil {
				var st string
				if mp3player.Playing {
					st = "play"
				}
				responder.WriteResponse(st)
			} else {
				responder.WriteResponse("none")
			}
		case 'u':
			if mp3player != nil {
				msg = mp3player.URL
			} else {
				msg = "none"
			}
			responder.WriteResponse(msg)
		case 'e':
			if mp3player != nil && mp3player.Err != nil {
				msg = mp3player.Err.Error()
			} else {
				msg = "none"
			}
			responder.WriteResponse(msg)
		case 'p':
			if mp3player != nil && mp3player.Playing {
				mp3player.Close()
			}
			// le paramètre commence en position 1
			url := buf[1:]
			var err error
			if mp3player, err = NewMP3player(url); err != nil {
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
