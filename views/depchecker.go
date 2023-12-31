package depchecker

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/config"
	"github.com/crbroughton/go-backstop/constants"
	"github.com/crbroughton/go-backstop/docker"
	"github.com/crbroughton/go-backstop/styles"
	"github.com/crbroughton/go-backstop/utils"
	"github.com/muesli/reflow/indent"
)

type dockerNotInstalled time.Duration
type dockerInstalled bool
type DependenciesInstalled bool

type result struct {
	emoji string
}

type Model struct {
	Spinner     spinner.Model
	result      result
	hasBackstop bool
	hasDocker   bool
}

func New() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = styles.SpinnerStyle
	return Model{
		Spinner: s,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Spinner.Tick, m.checkDocker, docker.CheckForImage)
}

func (m *Model) checkDocker() tea.Msg {
	pause := time.Duration(time.Second)
	time.Sleep(pause)
	_, err := exec.LookPath("docker")

	if utils.IsError(err) {
		m.hasDocker = false
		return dockerNotInstalled(pause)
	}

	config.CreateJSON()

	return dockerInstalled(true)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			return m, tea.Quit
		default:
			return m, nil
		}

	case docker.BackstopImageInstalled:
		m.hasBackstop = true
		return m, nil

	case dockerInstalled:
		m.hasDocker = true
		return m, nil

	case dockerNotInstalled:
		m.hasDocker = false
		m.result = result{emoji: "❌"}
		return m, nil

	default:
		var cmd tea.Cmd
		if m.hasBackstop && m.hasDocker {
			return m, func() tea.Msg {
				pause := time.Duration(time.Second)
				time.Sleep(pause)
				config.SetDependencyCheck()
				return DependenciesInstalled(true)
			}
		}
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	var spinnerMsg = m.Spinner.View() + " Checking dependencies, press ESC to quit...\n\n"
	s := "\n" + spinnerMsg

	// docker
	var hasDocker = "\n" + "✅ Docker found!"
	var noDocker = "\n" + "❌ Docker not found, searching..."

	// backstop
	var hasBackstop = "\n" + "✅ Backstop Installed!"
	var noBackstop = "\n" + "❌ Backstop image not found"

	if !m.hasDocker && len(m.result.emoji) > 0 {
		s := "\n" + m.result.emoji + " Docker is not installed! :( \n\n"
		s += ("\n\nPress any key to exit\n")

		return indent.String(s, 1)
	}

	if m.hasBackstop {
		s += hasBackstop
	} else {
		s += noBackstop
	}

	if m.hasDocker {
		s += hasDocker
	} else {
		s += noDocker

	}

	if m.hasDocker && m.hasBackstop {
		s = hasBackstop + hasDocker + "\n" + "✅ Dependency requirements met!"
	}

	s += ("\n\nPress any key to exit\n")

	return indent.String(s, 1)
}
