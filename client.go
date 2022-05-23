package main

import (
	"fmt"
	"goradio/cli"
	"log"
	"os"
	"os/exec"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
)

func client() error {
	_, err := os.Stat(namedpipe)
	if os.IsNotExist(err) {
		log.Println("client: creating namedpipe")
		if err := syscall.Mkfifo(namedpipe, 0666); err != nil {
			return err
		}
		// launch server
		log.Println("client: launching server")
		cmd := exec.Command(os.Args[0], "-d")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Start(); err != nil {
			return err
		}
	}

	pipe, err := os.OpenFile(namedpipe, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		return err
	}
	defer pipe.Close()

	p := tea.NewProgram(cli.InitModel(pipe), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	// log.Print("client: ask server to quit")
	// pipe.Write([]byte("q"))
	// log.Print("client: ask server to play")
	// pipe.Write([]byte("phttps://stream.klassikradio.de/meditation/mp3-192/radiode/"))
	// log.Print("client: ask server to another stream")
	// pipe.Write([]byte("phttp://37.187.93.104:8097/stream/1/"))
	// time.Sleep(5 * time.Second)
	return nil
}
