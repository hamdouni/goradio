package com

import "log"

const requestnamedpipepath = "/tmp/goradiorequest.np"

type PipePlayer struct {
	request Pipe
}

func NewPipePlayer() (PipePlayer, bool, error) {
	var pp PipePlayer

	requestpipe, created, err := NewPipe(requestnamedpipepath)
	if err != nil {
		return pp, created, err
	}
	pp.request = requestpipe
	return pp, created, nil
}

func (p PipePlayer) Close() {
	p.request.Close()
}

func (p PipePlayer) Play(url string) {
	p.request.Write("p" + url + "\n")
}

func (p PipePlayer) Pause() {
	p.request.Write("z")
}

func (p PipePlayer) Quit() {
	p.request.Write("q")
}

func (p PipePlayer) Read() (msg string) {
	msg, err := p.request.Read()
	if err != nil {
		log.Printf("PipePlayer.GetRequest got warning: %s", err)
	}
	return msg
}
