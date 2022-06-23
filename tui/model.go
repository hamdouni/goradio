package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/jamesnetherton/m3u"

	"goradio/player"

	tea "github.com/charmbracelet/bubbletea"
)

type station struct {
	name string
	url  string
	idx  int
}

func (st station) Title() string       { return st.name }
func (st station) Description() string { return st.url }
func (st station) FilterValue() string { return st.name }

type model struct {
	stations list.Model
	current  station
	message  string
	spin     spinner.Model
	player   player.Player
}

func initProcess() error {
	return nil
}

func initModel(player player.Player, playlist m3u.Playlist) (m model) {
	items := []list.Item{}
	var st station
	for i, track := range playlist.Tracks {
		st = station{
			name: track.Name,
			url:  track.URI,
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
	m.player = player
	return m
}

func (m model) getIndex(url string) int {
	if url == "none" {
		return -1
	}
	for idx, it := range m.stations.Items() {
		st := it.(station)
		if st.url == url {
			return idx
		}
	}
	return -1
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
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)
		case "enter":
			m.current = m.selected()
			m.player.Play(m.current.url)
			st := m.player.Status()
			if st.Err != nil {
				m.message = st.Err.Error()
			} else {
				m.message = m.current.name
			}
			// cmds = append(cmds, load(m.current.uri))
		case "Q":
			m.player.Quit()
			cmds = append(cmds, tea.Quit)
		case "o":
			st := m.player.Status()
			idx := m.getIndex(st.URL)
			if idx > -1 {
				m.current.idx = idx
				m.stations.Select(idx)
			}
		case " ":
			m.player.Pause()
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
	head := m.spin.View() + "ZiK> " + m.message
	return "  " + head + "\n\n" +
		m.stations.View()
}
