package com

import (
	"fmt"
	"log"

	"goradio/player"
)

const (
	requestnamedpipepath  = "/tmp/goradiorequest.np"
	responsenamedpipepath = "/tmp/goradioresponse.np"
)

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

func (p PipePlayer) Play(url string) {
	p.WriteRequest("p" + url)
}

func (p PipePlayer) Pause() {
	p.WriteRequest("z")
}

func (p PipePlayer) Quit() {
	p.WriteRequest("q")
}

func (p PipePlayer) Status() (st player.Stat) {
	p.WriteRequest("t")
	st.Title = p.ReadResponse()
	p.WriteRequest("u")
	st.URL = p.ReadResponse()
	p.WriteRequest("e")
	res := p.ReadResponse()
	if res == "etimeout" {
		st.Err = fmt.Errorf("no music")
	} else if res != "none" {
		st.Err = fmt.Errorf("%s", res)
	} else {
		st.Err = nil
	}
	return st
}

func (p PipePlayer) Close() {
	p.request.Close()
	p.response.Close()
}

func (p PipePlayer) ReadRequest() (msg string) {
	msg, err := p.request.Read()
	if err != nil {
		log.Printf("PipePlayer.ReadRequest got warning: %s", err)
	}
	return msg
}

func (p PipePlayer) ReadResponse() (msg string) {
	msg, err := p.response.ReadOrTimeout()
	if err != nil {
		log.Printf("PipePlayer.ReadResponse got warning: %s", err)
	}
	return msg
}

func (p PipePlayer) WriteRequest(msg string) {
	if err := p.request.Write(msg); err != nil {
		log.Printf("PipePlayer.WriteRequest got err: %s", err)
	}
}

func (p PipePlayer) WriteResponse(msg string) {
	if err := p.response.Write(msg); err != nil {
		log.Printf("PipePlayer.WriteResponse got err: %s", err)
	}
}
