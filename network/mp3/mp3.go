package mp3

import (
	"fmt"
	"time"

	"goradio/shoutcast"

	gomp3 "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type MP3player struct {
	player  oto.Player
	dec     *gomp3.Decoder
	stream  *shoutcast.Stream
	context *oto.Context
}

var context *oto.Context

func New(url string, callback func(*shoutcast.Metadata)) (mp3 *MP3player, err error) {

	// open shoutcast stream
	stream, err := shoutcast.Open(url)
	if err != nil {
		return mp3, fmt.Errorf("could not open shoutcast stream at %s: %s", url, err)
	}

	// register a callback function to be called when song changes
	stream.MetadataCallbackFunc = callback

	decoder, err := gomp3.NewDecoder(stream)
	if err != nil {
		return mp3, fmt.Errorf("could not decode stream %s", err)
	}
	if context == nil {
		var ready chan struct{}
		context, ready, err = oto.NewContext(decoder.SampleRate(), 2, 2)
		if err != nil {
			return mp3, fmt.Errorf("could not get oto context %s", err)
		}
		<-ready
	}
	player := context.NewPlayer(decoder)
	player.(oto.BufferSizeSetter).SetBufferSize(15000)

	mp3 = &MP3player{
		dec:     decoder,
		player:  player,
		stream:  stream,
		context: context,
	}

	return mp3, nil
}

func (mp3 *MP3player) Pause() {
	if mp3.player.IsPlaying() {
		mp3.player.Pause()
	} else {
		mp3.player.Play()
	}
}

func (mp3 *MP3player) Close() {
	mp3.stream.Close()
	mp3.player.Close()
}

func (mp3 *MP3player) Play() (err error) {
	go func() {
		defer mp3.stream.Close()
		mp3.player.Play()
		for {
			time.Sleep(1 * time.Second)
		}
	}()
	return nil
}
