package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/config"
	"github.com/crbroughton/go-backstop/utils"
	mainmenu "github.com/crbroughton/go-backstop/views/mainMenu"
	"github.com/crbroughton/go-backstop/views/mainMenu/settings"
)

type sessionState int

const (
	mainMenu sessionState = iota
	settingsMenu
)

type MainModel struct {
	state        sessionState
	mainMenu     tea.Model
	settingsMenu tea.Model
	windowSize   tea.WindowSizeMsg
}

func main() {
	model := New()
	program := tea.NewProgram(model, tea.WithAltScreen())

	_, err := program.Run()

	if utils.IsError(err) {
		fmt.Println(err)
		os.Exit(1)
	}
}

func New() MainModel {
	return MainModel{
		state:        mainMenu,
		mainMenu:     mainmenu.New(),
		settingsMenu: settings.SettingsModel,
	}
}

func (m MainModel) Init() tea.Cmd {
	_, err := exec.LookPath("docker")

	if utils.IsError(err) {
		fmt.Println("Could not find docker; Please install docker")
		os.Exit(1)
	}

	config.CreateJSON()
	config.WriteDefaultConfiguration()

	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg // pass this along to the entry view so it uses the full window size when it's initialized
	case mainmenu.SettingsSelected:
		m.state = settingsMenu
	case settings.GoBackToMainMenu:
		m.state = mainMenu
	}

	switch m.state {
	case mainMenu:
		newMainMenu, newCmd := m.mainMenu.Update(msg)
		mainMenuModel, ok := newMainMenu.(mainmenu.Model)

		if !ok {
			panic("could not perform assertion on mainmenu model")
		}
		m.mainMenu = mainMenuModel
		cmd = newCmd
	case settingsMenu:
		newSettingsMenu, newCmd := m.settingsMenu.Update(msg)
		settingsMenumodel, ok := newSettingsMenu.(settings.Model)

		if !ok {
			panic("could not perform assertion on settingsmenu model")
		}
		m.settingsMenu = settingsMenumodel
		cmd = newCmd
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {
	case mainMenu:
		return m.mainMenu.View()
	case settingsMenu:
		return m.settingsMenu.View()
	default:
		return m.mainMenu.View()
	}
}
