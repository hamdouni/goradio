package server

import (
	"fmt"
	"io"
	"net/http"

	gomp3 "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

type MP3player struct {
	Playing bool
	URL     string
	Err     error

	dec     *gomp3.Decoder
	player  *oto.Player
	stream  *http.Response
	context *oto.Context
}

func NewMP3player(url string) (mp3 *MP3player, err error) {

	mp3 = new(MP3player)

	mp3.Playing = false
	mp3.URL = url

	mp3.stream, err = http.Get(url)
	if err != nil {
		return mp3, err
	}
	if mp3.stream.StatusCode < 200 || mp3.stream.StatusCode > 299 {
		mp3.stream.Body.Close()
		mp3.Err = fmt.Errorf("erreur http")
		return mp3, mp3.Err
	}

	mp3.dec, err = gomp3.NewDecoder(mp3.stream.Body)
	if err != nil {
		mp3.Err = fmt.Errorf("erreur decodeur: %s", err)
		return mp3, mp3.Err
	}

	if mp3.context, err = oto.NewContext(mp3.dec.SampleRate(), 2, 2, 16384); err != nil {
		mp3.Err = fmt.Errorf("erreur context: %s", err)
		return mp3, mp3.Err
	}
	mp3.player = mp3.context.NewPlayer()

	return mp3, nil
}

func (mp3 *MP3player) Close() {
	mp3.Playing = false
}

func (mp3 *MP3player) Play() error {

	if mp3.Playing {
		mp3.Err = fmt.Errorf("mp3 player already playing")
		return mp3.Err
	}

	go func() {
		defer func() {
			if err := mp3.stream.Body.Close(); err != nil {
				mp3.Err = fmt.Errorf("mp3 body close: %s", err)
			}
			if err := mp3.player.Close(); err != nil {
				mp3.Err = fmt.Errorf("mp3 player close: %s", err)
			}
			if err := mp3.context.Close(); err != nil {
				mp3.Err = fmt.Errorf("mp3 context close: %s", err)
			}
		}()
		mp3.Playing = true
		data := make([]byte, 512)
		for mp3.Playing {
			_, err := mp3.dec.Read(data)
			if err == io.EOF || err != nil {
				continue
			}
			mp3.Err = nil
			mp3.player.Write(data)
		}
		mp3.Playing = false
	}()

	return nil
}
