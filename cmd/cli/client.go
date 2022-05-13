package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "./cmd/srv/", "http://37.187.93.104:8097/stream/1/")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0) // c'est l'astuce pour laisser le daemon en background sans être killé par la fin de ce programme
}
