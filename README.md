# Goradio

Mp3 stream player in the terminal written in go as a proof of concept.
The music is played in background so you can quit the terminal.

## Usage

Build and run :
```sh
git clone https://github.com/hamdouni/goradio
cd goradio
go build 
./goradio
```

You can also go install it if you have configured your go path binaries in your PATH :
```sh
go install
```

The program reads an M3U playlist, named `musics.m3u`, with only mp3 streams (no sublist as .m3u or .pls). 

It renders the content in a list that you can interact with :

- ↑/k   : up 
- ↓/j   : down 
- /     : filter
- enter : play selected stream
- space : toggle pause/play
- o     : jump to actual played station
- q     : quit but let music in background (also ctrl+c)
- Q     : quit and stop music

## Cross-compile [DEPRECATED]

**Not working since using namedpipe to communicate between client and server**

@TODO:
- [ ] use http server and client for inter process communication
- [ ] remove namedpipe code
- [ ] cross-compile !

```sh
GOOS=windows go build -o goradio.exe cmd/*
```

## Credits

This application uses the work of opensource contributors. 

Thank you to all of them for bringing those awesome projects.

The sound part exists because of those great libraries :

* [oto](https://github.com/hajimehoshi/oto) : A low-level library to play sound on multiple platforms (Apache-2.0 License)
* [go-mp3](https://github.com/hajimehoshi/go-mp3) : An MP3 decoder in pure Go (Apache-2.0 License)
* [m3u](https://github.com/jamesnetherton/m3u) : A basic golang m3u playlist parser  (MIT License) 
* [shoutcast](github.com/romantomjak/shoutcast) : All the heavy lifting of the shoutcast protocole to retrieve stream title (MIT License)

The interface could not have been more simple without the work of Charm and their products :

* [bubbletea](https://github.com/charmbracelet/bubbletea) : A powerful little TUI framework (MIT License) 
* [bubbles](https://github.com/charmbracelet/bubbles) : TUI components for Bubble Tea (MIT License) 
