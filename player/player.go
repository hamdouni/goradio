package player

type Player interface {
	Play(url string)
	Pause()
	Status() string
	Quit()
}

type Responder interface {
	ReadRequest() string
	Write(msg string)
	Close()
}
