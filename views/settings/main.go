package settings

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type SettingsPageItem struct {
	title string
}

func (i SettingsPageItem) Title() string       { return i.title }
func (i SettingsPageItem) FilterValue() string { return i.title }

type SettingsPage struct {
	List  list.Model
	title string
}

func (m SettingsPage) Init() tea.Cmd {
	return nil
}

func (m SettingsPage) View() string {
	return "user settings"
}

func (m SettingsPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func Content() []list.Item {
	return []list.Item{
		SettingsPageItem{title: "Create user cookie"},
	}
}

var content = Content()

var delegate = list.NewDefaultDelegate()

var Modal = SettingsPage{
	List:  list.New(content, delegate, 0, 0),
	title: "SettingsPage",
}
