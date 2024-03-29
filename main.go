package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/adrg/xdg"
	"github.com/jamesnetherton/m3u"

	"goradio/com"
	"goradio/server"
	"goradio/tui"
)

func main() {
	daemon := flag.Bool("d", false, "daemon server mode")
	info := flag.Bool("i", false, "get information on played music")
	pause := flag.Bool("p", false, "pause/unpause music")
	flag.Parse()

	pipeplayer, created, err := com.NewPipePlayer()
	if err != nil {
		log.Fatalf("com.NewPipePlayer returns %v", err)
	}

	if *daemon {
		// mode server : on lit le namedpipe

		dlog, _ := os.OpenFile("/tmp/goradio-daemon.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		defer dlog.Close()
		log.SetOutput(dlog)
		if err = server.Run(pipeplayer); err != nil {
			log.Fatalf("server returns %v", err)
		}
		os.Exit(0)
	}

	if created && (*info || *pause) {
		fmt.Println("no music")
		return
	}

	if *pause {
		pipeplayer.Pause()
		return
	}

	musicFile, err := xdg.DataFile("goradio/musics.m3u")
	if err != nil {
		log.Fatalf("erreur %s\n", err)
	}
	playlist, err := m3u.Parse(musicFile)
	if err != nil {
		log.Fatal(err)
	}

	if *info {
		status := pipeplayer.Status()
		if status.Err != nil {
			m := 20
			if len(status.Err.Error()) < m {
				m = len(status.Err.Error())
			}
			fmt.Printf("%s", status.Err.Error()[0:m])
			return
		} else if status.Title != "" {
			fmt.Print(status.Title)
			return
		} else if status.URL == "" || status.URL[0:4] != "http" {
			fmt.Print("no music")
			return
		}
		for _, track := range playlist.Tracks {
			if track.URI == status.URL {
				fmt.Println(track.Name)
				return
			}
		}
		fmt.Printf("track not found: %s", status.URL)
		return
	}

	// mode cli
	clog, _ := os.OpenFile("/tmp/goradio-client.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	defer clog.Close()
	log.SetOutput(clog)
	status := pipeplayer.Status()
	log.Printf("debug: %v\n", status)
	if status.URL == "etimeout" {
		// launch server
		log.Println("client: launching server")
		cmd := exec.Command(os.Args[0], "-d")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Start(); err != nil {
			log.Fatalf("launching server returns %v", err)
		}
	}
	if err := tui.Run(pipeplayer, playlist); err != nil {
		log.Fatalf("client returns %v", err)
	}
}
