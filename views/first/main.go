package first

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/styles"
)

func (m MainPage) Init() tea.Cmd {
	return nil
}

type MainPageItem struct {
	title string
}

func (i MainPageItem) Title() string       { return i.title }
func (i MainPageItem) FilterValue() string { return i.title }

type MainPage struct {
	List  list.Model
	title string
}

func (m MainPage) View() string {
	return styles.AppStyle.Render(m.List.View())
}

func (m MainPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func Content() []list.Item {
	return []list.Item{
		MainPageItem{
			title: "test",
		},
		MainPageItem{
			title: "test2",
		},
	}
}

var content = Content()

var delegate = list.NewDefaultDelegate()

var Modal = MainPage{
	List:  list.New(content, delegate, 0, 0),
	title: "masterPage",
}
