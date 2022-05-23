package com

import (
	"os"
	"syscall"
)

type Piper struct {
	namedpipe *os.File
}

const namedpipepath = "/tmp/goradio.np"

func New() (*Piper, error) {
	var p Piper
	_, err := os.Stat(namedpipepath)
	if os.IsNotExist(err) {
		if err := syscall.Mkfifo(namedpipepath, 0666); err != nil {
			return nil, err
		}
	}
	np, err := os.OpenFile(namedpipepath, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		return nil, err
	}
	p.namedpipe = np
	return &p, nil
}

func (p *Piper) ping() bool {
	msg := []byte{'g'}
	p.namedpipe.Write(msg)
	return true
}

func (p *Piper) play(url string) error {
	return nil
}

func (p *Piper) pause() error {
	return nil
}

func (p *Piper) quit() error {
	p.namedpipe.Close()
	os.Remove(namedpipepath)
	return nil
}
