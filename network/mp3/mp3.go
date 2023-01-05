package mp3

import (
	"fmt"
	"net/http"
	"time"

	gomp3 "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type MP3player struct {
	player  oto.Player
	dec     *gomp3.Decoder
	stream  *http.Response
	context *oto.Context
}

func New(url string) (mp3 *MP3player, err error) {
	stream, err := http.Get(url)
	if err != nil {
		return mp3, fmt.Errorf("could not get url %s", err)
	}
	if stream.StatusCode < 200 || stream.StatusCode > 299 {
		stream.Body.Close()
		return mp3, fmt.Errorf("erreur http")
	}
	decoder, err := gomp3.NewDecoder(stream.Body)
	if err != nil {
		return mp3, fmt.Errorf("could not decode stream %s", err)
	}
	context, ready, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		return mp3, fmt.Errorf("could not get oto context %s", err)
	}
	<-ready
	player := context.NewPlayer(decoder)

	mp3 = &MP3player{
		dec:     decoder,
		player:  player,
		stream:  stream,
		context: context,
	}

	return mp3, nil
}

func (mp3 *MP3player) Close() {
	mp3.player.Close()
}

func (mp3 *MP3player) Play() (err error) {
	go func() {
		defer mp3.stream.Body.Close()
		mp3.player.Play()
		for mp3.player.IsPlaying() {
			time.Sleep(1 * time.Second)
		}
	}()
	return nil
}
