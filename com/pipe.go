package com

import (
	"os"
	"syscall"
)

type Pipe struct {
	pipefile *os.File
}

func NewPipe(pipepath string) (p Pipe, created bool, err error) {

	np, created, err := openOrCreatePipe(pipepath)
	if err != nil {
		return p, created, err
	}

	p.pipefile = np

	return p, created, nil
}

func (p *Pipe) Close() error {
	err := p.pipefile.Close()
	if err != nil {
		return err
	}
	return os.Remove(p.pipefile.Name())
}

func (p *Pipe) Read() (string, error) {
	buf := make([]byte, 255)
	_, err := p.pipefile.Read(buf)
	if err != nil {
		return "", err
	}
	// extract until \n or null
	var result string
	for _, v := range buf {
		if v == 0 || v == '\n' {
			break
		}
		result += string(v)
	}
	return result, nil
}

func (p *Pipe) Write(msg string) error {
	_, err := p.pipefile.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func openOrCreatePipe(pipepath string) (*os.File, bool, error) {
	created := false
	_, err := os.Stat(pipepath)
	if os.IsNotExist(err) {
		created = true
		if err := syscall.Mkfifo(pipepath, 0666); err != nil {
			return nil, created, err
		}
	}
	np, err := os.OpenFile(pipepath, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		return nil, created, err
	}
	return np, created, nil
}
