package mp3

import (
	"fmt"
	"io"
	"log"
	"net/http"

	gomp3 "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

type MP3player struct {
	Playing bool
	Paused  bool

	dec     *gomp3.Decoder
	player  *oto.Player
	stream  *http.Response
	context *oto.Context
	data    []byte
}

func New(url string) (mp3 *MP3player, err error) {

	mp3 = new(MP3player)

	mp3.data = make([]byte, 512)
	mp3.Playing = false

	mp3.stream, err = http.Get(url)
	if err != nil {
		return mp3, err
	}
	if mp3.stream.StatusCode < 200 || mp3.stream.StatusCode > 299 {
		mp3.stream.Body.Close()
		return mp3, fmt.Errorf("erreur http")
	}

	mp3.dec, err = gomp3.NewDecoder(mp3.stream.Body)
	if err != nil {
		return mp3, err
	}

	if mp3.context, err = oto.NewContext(mp3.dec.SampleRate(), 2, 2, 16384); err != nil {
		return mp3, err
	}
	mp3.player = mp3.context.NewPlayer()

	return mp3, nil
}

func (mp3 *MP3player) Close() {
	mp3.Playing = false
}

func (mp3 *MP3player) Play() (err error) {

	if mp3.Playing {
		return fmt.Errorf("mp3 player already playing")
	}

	go func() {
		defer func() {
			if err := mp3.stream.Body.Close(); err != nil {
				log.Printf("body close: %s", err)
			}
			if err := mp3.player.Close(); err != nil {
				log.Printf("player close: %s", err)
			}
			if err := mp3.context.Close(); err != nil {
				log.Printf("context close: %s", err)
			}
		}()
		mp3.Playing = true
		for mp3.Playing {
			_, err = mp3.dec.Read(mp3.data)
			if err == io.EOF || err != nil {
				mp3.Playing = false
			}
			if mp3.Playing && !mp3.Paused {
				mp3.player.Write(mp3.data)
			}
		}
	}()

	return nil
}
