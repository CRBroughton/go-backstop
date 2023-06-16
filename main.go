package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/iterator"
	"github.com/crbroughton/go-backstop/styles"
	"github.com/crbroughton/go-backstop/utils"
	master "github.com/crbroughton/go-backstop/views/first"
	"github.com/crbroughton/go-backstop/views/second"
)

type item struct {
	title   string
	desc    string
	command string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list    list.Model
	focused iterator.Status
	loading bool
}

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
		item{title: "Run tests", desc: "Runs all stored tests", command: ""},
		item{title: "Create new test", desc: "Create a new test for your site"},
		item{title: "Change user settings", desc: "Update your personal settings"},
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				response := utils.RunCommand(i.command, &m.loading)
				m.loading = response
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := styles.DocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
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

	delegate.Styles.SelectedTitle = styles.SelectedItem
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedTitle.Copy()

	m := model{list: list.New(items, delegate, 0, 0)}

	m.list.Styles.Title = styles.TitleStyle
	m.list.Title = "Go, Backstop!"

	program := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
