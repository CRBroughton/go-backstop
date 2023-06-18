package mainmenu

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/iterator"
	"github.com/crbroughton/go-backstop/styles"
)

type Model struct {
	list list.Model
}

type item struct {
	title string
	desc  string
	ID    iterator.Page
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func New() tea.Model {
	delegate := list.NewDefaultDelegate()

	model := Model{list: list.New(
		Content(),
		delegate,
		0,
		0,
	)}

	model.list.Title = "Settings"
	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := styles.AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch {
		default:
			m.list, cmd = m.list.Update(msg)
		}
		cmds = append(cmds, cmd)
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return styles.AppStyle.Render(m.list.View())
}
func Content() []list.Item {
	return []list.Item{
		item{title: "Create user cookie", desc: "test"},
		item{title: "Create user cookie", desc: "test"},
	}
}
