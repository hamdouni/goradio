package com

type Communicator interface {
	play(string) error
	pause() error
	quit() error
}
