package com

import (
	"os"
	"syscall"
	"time"
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

// Read reads the pipe until any data is available
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

// ReadOrTimeout reads the pipe but timeout after some time
func (p *Pipe) ReadOrTimeout() (string, error) {

	const timeout = 300 * time.Millisecond

	ch := make(chan string, 1)

	go func() {
		buf := make([]byte, 255)
		_, err := p.pipefile.Read(buf)
		if err != nil {
			ch <- "e" + err.Error()
		}
		// extract until \n or null
		var res string
		for _, v := range buf {
			if v == 0 || v == '\n' {
				ch <- res // send it
				res = ""  // ready for next token
			}
			res += string(v)
		}
	}()

	var result string
	select {
	case result = <-ch:
	case <-time.After(timeout):
		result = "etimeout"
	}

	return result, nil
}

func (p *Pipe) Write(msg string) error {
	_, err := p.pipefile.WriteString(msg + "\n")
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
