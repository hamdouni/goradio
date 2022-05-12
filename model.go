package main

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jamesnetherton/m3u"
)

type station struct {
	name string
	uri  string
	idx  int
}

func (st station) Title() string       { return st.name }
func (st station) Description() string { return st.uri }
func (st station) FilterValue() string { return st.name }

type model struct {
	stations list.Model
	current  station
	message  string
	spin     spinner.Model
}

func initProcess() error {
	return nil
}

func initModel() (m model) {
	playlist, err := m3u.Parse("musics.m3u")
	if err != nil {
		log.Fatal(err)
	}
	items := []list.Item{}
	var st station
	for i, track := range playlist.Tracks {
		st = station{
			name: track.Name,
			uri:  track.URI,
			idx:  i,
		}
		items = append(items, st)
	}
	m.stations = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.stations.DisableQuitKeybindings()
	m.message = "List initialized..."
	s := spinner.New()
	s.Spinner = spinner.Dot
	m.spin = s
	return m
}

func (m model) Init() tea.Cmd {
	return m.spin.Tick
}

func (m model) selected() station {
	st, ok := m.stations.SelectedItem().(station)
	if !ok {
		return station{}
	}
	return st
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case playerError:
		m.message = msg.err.Error()
	case playerStarted:
		m.message = msg.status
	case playerLoaded:
		cmds = append(cmds, play(m.current.name))
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)
		case "enter":
			m.current = m.selected()
			cmds = append(cmds, load(m.current.uri))
		case "o":
			m.stations.Select(m.current.idx)
		case " ":
			if player != nil && player.Playing {
				player.Close()
				m.message = "pause " + m.current.name
			} else if m.current.uri != "" {
				cmds = append(cmds, load(m.current.uri))
			}
		}
	case tea.WindowSizeMsg:
		m.stations.SetSize(msg.Width, msg.Height-3)
	}

	var cmd tea.Cmd

	m.stations, cmd = m.stations.Update(msg)
	cmds = append(cmds, cmd)

	m.spin, cmd = m.spin.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() (s string) {
	var head = "ZiK> " + m.message
	if player != nil && player.Playing {
		head = m.spin.View() + head
	}
	return "  " + head + "\n\n" +
		m.stations.View()
}
