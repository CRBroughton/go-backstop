package depchecker

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/crbroughton/go-backstop/config"
	"github.com/crbroughton/go-backstop/constants"
	"github.com/crbroughton/go-backstop/docker"
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
	hasDeps     bool
	hasBackstop bool
	hasDocker   bool
}

func New() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
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
				return DependenciesInstalled(true)
			}
		}
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	s := "\n" +
		m.Spinner.View() + " Checking dependencies, press ESC to quit...\n\n"

	if !m.hasDocker && len(m.result.emoji) > 0 {
		s := "\n" + m.result.emoji + " Docker is not installed! :( \n\n"
		s += ("\n\nPress any key to exit\n")

		return indent.String(s, 1)
	}

	if m.hasBackstop {
		s = s + "\n" + "✅ Backstop Installed!"
	} else {
		s = s + "\n" + "❌ Backstop image not found"
	}

	if m.hasDocker {
		s = s + "\n" + m.result.emoji + "✅ Docker found!"
	} else {
		s = s + "\n" + "❌ Docker not found, searching..."

	}

	if m.hasDocker && m.hasBackstop {
		s = s + "\n" + "✅ Dependency requirements met!"
	}

	s += ("\n\nPress any key to exit\n")

	return indent.String(s, 1)
}
