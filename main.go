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
	"github.com/crbroughton/go-backstop/views/mainMenu/settings"
	"github.com/crbroughton/go-backstop/views/mainMenu/settings/cookies"
	"github.com/crbroughton/go-backstop/views/mainMenu/settings/viewport"
)

type sessionState int

const (
	depChecker sessionState = iota
	mainMenu
	settingsMenu
	cookieMenu
	viewportMenu
)

type MainModel struct {
	state        sessionState
	depChecker   tea.Model
	mainMenu     tea.Model
	settingsMenu tea.Model
	cookieMenu   tea.Model
	viewportMenu tea.Model
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
		depChecker:   depchecker.New(),
		mainMenu:     mainmenu.New(),
		settingsMenu: settings.New(),
		cookieMenu:   cookies.New(),
		viewportMenu: viewport.New(),
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
	case mainmenu.SettingsSelected:
		m.state = settingsMenu
	case settings.GoBackToMainMenu:
		m.state = mainMenu
	case settings.GoToViewPort:
		m.state = viewportMenu
	case settings.GoToCookie:
		m.state = cookieMenu
	case viewport.GoBackToSettingsMenu:
		m.state = settingsMenu
	case cookies.GoBackToSettingsMenu:
		m.state = settingsMenu
	case spinner.TickMsg:
		m.depChecker, cmd = m.depChecker.Update(msg)
		cmds = append(cmds, cmd)
	}

	switch m.state {
	case depChecker:
		m.depChecker, cmd = m.depChecker.Update(msg)
		cmds = append(cmds, cmd)
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
	case cookieMenu:
		newCookieMenu, newCmd := m.cookieMenu.Update(msg)
		cookieMenuModel, ok := newCookieMenu.(cookies.Model)

		if !ok {
			panic("could not perform assertion on cookiemenu model")
		}
		m.cookieMenu = cookieMenuModel
		cmd = newCmd
	case viewportMenu:
		newViewportMenu, newCmd := m.viewportMenu.Update(msg)
		viewportMenuModel, ok := newViewportMenu.(viewport.Model)

		if !ok {
			panic("could not perform assertion on viewportmenu model")
		}
		m.viewportMenu = viewportMenuModel
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
	case cookieMenu:
		return m.cookieMenu.View()
	case viewportMenu:
		return m.viewportMenu.View()
	default:
		return m.depChecker.View()
	}
}
