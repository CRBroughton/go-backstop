package mainmenu

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/styles"
	"github.com/crbroughton/go-backstop/utils"
)

type (
	menuItem         int
	SettingsSelected bool
)

type Model struct {
	list     list.Model
	selected menuItem
}

const (
	runTests menuItem = iota
	createNewTest
	createReferenceImages
	settingsPage
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
				switch m.selected {
				case settingsPage:
					return m, func() tea.Msg {
						return SettingsSelected(true)
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
		item{title: "Run tests", desc: "Runs all stored tests ", ID: runTests},
		item{title: "Create new test", desc: "Create a new test", ID: createNewTest},
		item{title: "Create reference images", desc: "New ref images", ID: createReferenceImages},
		item{title: "Settings Page", desc: "Update your personal settings", ID: settingsPage},
	}
}

func (model *Model) setView(id menuItem) {
	model.selected = id

	switch id {
	case createNewTest:
	case runTests:
		runBackstopCommand("test")
	case createReferenceImages:
		runBackstopCommand("reference")

	}
}

func runBackstopCommand(command string) {
	workingDIR, err := os.Getwd()
	if utils.IsError(err) {
		log.Fatal(err)
	}

	args := []string{
		"run",
		"--rm",
		"-v",
		workingDIR + ":/src",
		"backstopjs/backstopjs",
		command,
		"--config=.settings/backstop.config.js",
	}
	err = utils.RunCommand("docker", args...)

	if utils.IsError(err) {
		log.Fatal(err)
	}
}
