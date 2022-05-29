package player

type Player interface {
	Play(url string)
	Pause()
	Quit()
}

type Responder interface {
	Read() string
	Close()
}
