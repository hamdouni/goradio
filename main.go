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
		if err := server.Run(pipeplayer); err != nil {
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

  musicFile,err := xdg.DataFile("goradio/musics.m3u")
  if err != nil {
    log.Fatalf("erreur %s\n",err)
  }
	playlist, err := m3u.Parse(musicFile)
	if err != nil {
		log.Fatal(err)
	}

	if *info {
		status := pipeplayer.Status()
		var name string
		if status != "" {
			for _, track := range playlist.Tracks {
				if track.URI == status {
					name = track.Name
				}
			}
		}
		fmt.Println(name)
		return
	}

	// mode cli
	if created {
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
