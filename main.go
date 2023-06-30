package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/config"
	"github.com/crbroughton/go-backstop/utils"
	depchecker "github.com/crbroughton/go-backstop/views"
	mainmenu "github.com/crbroughton/go-backstop/views/mainMenu"
	"github.com/crbroughton/go-backstop/views/mainMenu/createTests"
	"github.com/crbroughton/go-backstop/views/mainMenu/createTests/resultsTable"
	"github.com/crbroughton/go-backstop/views/mainMenu/settings"
	"github.com/crbroughton/go-backstop/views/mainMenu/settings/cookies"
	"github.com/crbroughton/go-backstop/views/mainMenu/settings/viewport"
)

type sessionState int

const (
	depChecker sessionState = iota
	mainMenu
	resultsTableMenu
	createTestMenu
	settingsMenu
	cookieMenu
	viewportMenu
)

type MainModel struct {
	state            sessionState
	depChecker       tea.Model
	mainMenu         tea.Model
	resultsTableMenu tea.Model
	createTestMenu   tea.Model
	settingsMenu     tea.Model
	cookieMenu       tea.Model
	viewportMenu     tea.Model
	windowSize       tea.WindowSizeMsg
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
		state:      depChecker,
		depChecker: depchecker.New(),
	}
}

func (m MainModel) Init() tea.Cmd {
	_, err := exec.LookPath("docker")

	if utils.IsError(err) {
		fmt.Println("Could not find docker; Please install docker")
		os.Exit(1)
	}

	config.CreateJSON()

	_, err = os.Stat("backstop_data")

	if errors.Is(err, os.ErrNotExist) {
		config.RunBackstopCommand("init", false)
	}

	_, err = os.Stat("backstop.json")

	if err == nil {
		err = os.Remove("backstop.json")
		if utils.IsError(err) {
			log.Fatal(err)
		}
	}

	return tea.Batch(m.depChecker.Init())
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg // pass this along to the entry view so it uses the full window size when it's initialized
	case depchecker.DependenciesInstalled:
		m.state = mainMenu
		m.mainMenu = mainmenu.New()
		m.mainMenu, cmd = m.mainMenu.Update(m.windowSize)
	case mainmenu.RunTestsSelected:
		m.state = resultsTableMenu
		m.resultsTableMenu = resultsTable.New()
	case mainmenu.CreatedNewTestSelected:
		m.state = createTestMenu
		m.createTestMenu = createTests.New()
	case createTests.GoBackToMainMenu:
		m.state = mainMenu
	case resultsTable.GoBackToSettingsMenu:
		m.state = mainMenu
		m.resultsTableMenu, cmd = m.resultsTableMenu.Update(m.windowSize)
	case mainmenu.SettingsSelected:
		m.state = settingsMenu
		m.settingsMenu = settings.New()
		m.settingsMenu, cmd = m.settingsMenu.Update(m.windowSize)
	case settings.GoBackToMainMenu:
		m.state = mainMenu
		m.mainMenu, cmd = m.mainMenu.Update(m.windowSize)
	case settings.GoToViewPort:
		m.state = viewportMenu
		m.viewportMenu = viewport.New()
		m.viewportMenu, cmd = m.viewportMenu.Update(m.windowSize)
	case settings.GoToCookie:
		m.state = cookieMenu
		m.cookieMenu = cookies.New()
		m.cookieMenu, cmd = m.cookieMenu.Update(m.windowSize)
	case viewport.GoBackToSettingsMenu:
		m.state = settingsMenu
		m.settingsMenu, cmd = m.settingsMenu.Update(m.windowSize)
	case cookies.GoBackToSettingsMenu:
		m.state = settingsMenu
		m.settingsMenu, cmd = m.settingsMenu.Update(m.windowSize)
	case spinner.TickMsg:
		m.depChecker, cmd = m.depChecker.Update(msg)
		cmds = append(cmds, cmd)
	}

	switch m.state {
	case depChecker:
		m.depChecker, cmd = m.depChecker.Update(msg)
		cmds = append(cmds, cmd)
	case mainMenu:
		m.mainMenu, cmd = m.mainMenu.Update(msg)
		cmds = append(cmds, cmd)
	case resultsTableMenu:
		m.resultsTableMenu, cmd = m.resultsTableMenu.Update(msg)
		cmds = append(cmds, cmd)
	case createTestMenu:
		m.createTestMenu, cmd = m.createTestMenu.Update(msg)
		cmds = append(cmds, cmd)
	case settingsMenu:
		m.settingsMenu, cmd = m.settingsMenu.Update(msg)
		cmds = append(cmds, cmd)

	case cookieMenu:
		m.cookieMenu, cmd = m.cookieMenu.Update(msg)
		cmds = append(cmds, cmd)
	case viewportMenu:
		m.viewportMenu, cmd = m.viewportMenu.Update(msg)
		cmds = append(cmds, cmd)
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {
	case mainMenu:
		return m.mainMenu.View()
	case resultsTableMenu:
		return m.resultsTableMenu.View()
	case createTestMenu:
		return m.createTestMenu.View()
	case settingsMenu:
		return m.settingsMenu.View()
	case cookieMenu:
		return m.cookieMenu.View()
	case viewportMenu:
		return m.viewportMenu.View()
	default:
		return m.depChecker.View()
	}
}
