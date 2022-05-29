package com

import "log"

const requestnamedpipepath = "/tmp/goradiorequest.np"
const responsenamedpipepath = "/tmp/goradioresponse.np"

type PipePlayer struct {
	request  Pipe
	response Pipe
}

func NewPipePlayer() (PipePlayer, bool, error) {
	var pp PipePlayer

	responsepipe, created, err := NewPipe(responsenamedpipepath)
	if err != nil {
		return pp, false, err
	}
	pp.response = responsepipe

	requestpipe, _, err := NewPipe(requestnamedpipepath)
	if err != nil {
		return pp, false, err
	}
	pp.request = requestpipe

	return pp, created, nil
}

func (p PipePlayer) Close() {
	p.request.Close()
	p.response.Close()
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

func (p PipePlayer) Status() (status string) {
	p.request.Write("s")
	return p.ReadResponse()
}

func (p PipePlayer) ReadRequest() (msg string) {
	msg, err := p.request.Read()
	if err != nil {
		log.Printf("PipePlayer.ReadRequest got warning: %s", err)
	}
	return msg
}

func (p PipePlayer) ReadResponse() (msg string) {
	msg, err := p.response.Read()
	if err != nil {
		log.Printf("PipePlayer.ReadResponse got warning: %s", err)
	}
	return msg
}

func (p PipePlayer) Write(msg string) {
	p.response.Write(msg)
}
