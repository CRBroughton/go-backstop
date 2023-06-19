package mainmenu

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/styles"
	"github.com/crbroughton/go-backstop/views/mainMenu/settings"
)

type Model struct {
	list     list.Model
	selected menuItem
}

type menuItem int

const (
	runTests menuItem = iota
	createNewTest
	settingsPage
)

type SettingsSelected bool

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
		0,
		0,
	)}

	model.list.Title = "Main menu"
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
		switch msg.String() {
		case "enter":
			item, ok := m.list.SelectedItem().(item)
			if !ok {
				log.Fatal("v ...any")
			}
			if ok {
				m.setView(item.ID)
				return m, func() tea.Msg {
					return SettingsSelected(true)
				}
			}
		}
		switch {
		// case key.Matches(msg, constants.Keymap.Enter):
		// 	cmd = selectProjectCmd(m.getSelectedMenuItem())
		default:
			m.list, cmd = m.list.Update(msg)
		}
		cmds = append(cmds, cmd)
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	switch m.selected {
	case settingsPage:
		return styles.AppStyle.Render(settings.SettingsModel.View())
	default:
		return styles.AppStyle.Render(m.list.View())
	}
}

func Content() []list.Item {
	return []list.Item{
		item{title: "Run tests", desc: "Runs all stored tests ", ID: runTests},
		item{title: "Create new test", desc: "Create a new test", ID: createNewTest},
		item{title: "Settings Page", desc: "Update your personal settings", ID: settingsPage},
	}
}

func (model *Model) setView(id menuItem) {
	model.selected = id
}

// func (m Model) getSelectedMenuItem() int {
// 	menuItems := m.list.Items()
// 	activeItem := menuItems[m.list.Index()]
// 	return activeItem.(item).ID
// }

// func selectProjectCmd(ActiveProjectID int) tea.Cmd {
// 	return func() tea.Msg {
// 		return SelectMsg{ActiveProjectID: ActiveProjectID}
// 	}
// }
