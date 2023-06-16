package second

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/styles"
)

func (m SecondPage) Init() tea.Cmd {
	return nil
}

type SecondPageItem struct {
	title string
}

func (i SecondPageItem) Title() string       { return i.title }
func (i SecondPageItem) FilterValue() string { return i.title }

type SecondPage struct {
	List  list.Model
	title string
}

func (m SecondPage) View() string {
	return styles.AppStyle.Render(m.List.View())
}

func (m SecondPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		SecondPageItem{
			title: "here",
		},
	}
}

var content = Content()

var delegate = list.NewDefaultDelegate()

var Modal = SecondPage{
	List:  list.New(content, delegate, 0, 0),
	title: "secondPage",
}
