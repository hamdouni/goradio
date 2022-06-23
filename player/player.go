package player

type Stat struct {
	Playing bool
	Err     error
	URL     string
}

type Player interface {
	Play(url string)
	Pause()
	Status() Stat
	Quit()
}

type Responder interface {
	ReadRequest() string
	Write(msg string)
	Close()
}
