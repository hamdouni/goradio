# Goradio

Mp3 stream player in the terminal written in go as a proof of concept.

## Usage

Build and run :
```sh
git clone https://github.com/hamdouni/goradio
cd goradio
go build -o goradio cmd/* 
./goradio
```

The program reads an M3U playlist, named `musics.m3u`, with only mp3 streams (no sublist as .m3u or .pls). 

It renders the content in a list that you can interact with :

-  ↑/k : up 
-  ↓/j : down 
-  /   : filter
- enter: play selected stream
- space: toggle pause/play
- q/ctrl+c: quit

## Cross-compile

```sh
GOOS=windows go build -o goradio.exe cmd/*
```

## Credits

This application uses the work of opensource contributors. 

Thank you to all of them for bringing those awesome projects.

The sound part exists because of those great libraries :

[oto](https://github.com/hajimehoshi/oto) : A low-level library to play sound on multiple platforms (Apache-2.0 License)

[go-mp3](https://github.com/hajimehoshi/go-mp3) : An MP3 decoder in pure Go (Apache-2.0 License)

[m3u](https://github.com/jamesnetherton/m3u) : A basic golang m3u playlist parser  (MIT License) 

The interface could not have been more simple unless the work of Charm and their products :

[bubbletea](https://github.com/charmbracelet/bubbletea) : A powerful little TUI framework (MIT License) 

[bubbles](https://github.com/charmbracelet/bubbles) : TUI components for Bubble Tea (MIT License) 
