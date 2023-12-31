package settings

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/constants"
	"github.com/crbroughton/go-backstop/styles"
)

type (
	menuItem         int
	GoToCookie       bool
	GoToViewPort     bool
	GoBackToMainMenu bool
)

type Model struct {
	list     list.Model
	selected menuItem
}

const (
	first menuItem = iota
	createCookie
	viewPort
	mainMenu
)

type item struct {
	title string
	desc  string
	ID    menuItem
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func New() tea.Model {
	delegate := list.NewDefaultDelegate()

	model := Model{list: list.New(
		Content(),
		delegate,
		50,
		20,
	)}

	model.list.Title = "Settings menu"
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
		case key.Matches(msg, constants.Keymap.Enter):
			item, ok := m.list.SelectedItem().(item)
			if !ok {
				log.Fatal("Something went wrong when selecting the item in the list")
			}
			if ok {
				m.setView(item.ID)
				switch m.selected {
				case createCookie:
					return m, func() tea.Msg {
						return GoToCookie(true)
					}
				case mainMenu:
					return m, func() tea.Msg {
						return GoBackToMainMenu(true)
					}
				case viewPort:
					return m, func() tea.Msg {
						return GoToViewPort(true)
					}
				}
			}
		}
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
		item{title: "Create user cookie", desc: "Create new cookie", ID: createCookie},
		item{title: "Add a viewport", desc: "New viewport", ID: viewPort},
		item{title: "Go back to main menu", ID: mainMenu},
	}
}

func (model *Model) setView(id menuItem) {
	model.selected = id
}
