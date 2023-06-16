package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/iterator"
	"github.com/crbroughton/go-backstop/styles"
	master "github.com/crbroughton/go-backstop/views/first"
	"github.com/crbroughton/go-backstop/views/second"
)

type model struct {
	list    list.Model
	focused iterator.Status
}

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) FilterValue() string { return i.title }

func (m model) Init() tea.Cmd {
	return nil
}

func (model *model) Next() {
	if model.focused < 1 {
		model.focused++
		second.Content()
	} else {
		model.focused = 0
	}
}

func items() []list.Item {
	return []list.Item{
		item{title: "test"},
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			m.Next()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.focused {
	case iterator.Root:
		return styles.AppStyle.Render(m.list.View())
	case iterator.MainPage:
		return master.Modal.View()
	case iterator.SecondPage:
		return second.Modal.View()
	default:
		return "unknown state"
	}
}

func main() {
	items := items()
	delegate := list.NewDefaultDelegate()
	m := model{list: list.New(items, delegate, 0, 0)}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error %v", err)
		os.Exit(1)
	}
}
