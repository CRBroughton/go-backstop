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
	"github.com/crbroughton/go-backstop/utils"
	"github.com/muesli/reflow/indent"
)

type dockerNotInstalled time.Duration
type dependenciesInstalled time.Duration

type result struct {
	duration time.Duration
	emoji    string
}

type Model struct {
	Spinner spinner.Model
	result  result
	hasDeps bool
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
	return tea.Batch(m.Spinner.Tick, m.checkDocker)
}

func (m *Model) checkDocker() tea.Msg {
	pause := time.Duration(time.Second)
	time.Sleep(pause)
	_, err := exec.LookPath("docker")

	if utils.IsError(err) {
		m.hasDeps = false
		return dockerNotInstalled(pause)

	}

	config.CreateJSON()
	m.hasDeps = true
	return dependenciesInstalled(pause)
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

	case dependenciesInstalled:
		d := time.Duration(msg)
		m.result = result{emoji: "✅", duration: d}
		m.hasDeps = true
		return m, nil

	case dockerNotInstalled:
		d := time.Duration(msg)
		m.result = result{emoji: "❌", duration: d}
		m.hasDeps = false
		return m, nil

	default:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	s := "\n" +
		m.Spinner.View() + " Checking dependencies, press ESC to quit...\n\n"

	if !m.hasDeps && len(m.result.emoji) > 0 {
		s := "\n" + m.result.emoji + " Docker is not installed! :( \n\n"
		s += ("\nPress any key to exit\n")

		return indent.String(s, 1)
	}

	if m.hasDeps && len(m.result.emoji) > 0 {
		s := "\n" + m.result.emoji + " Dependency requirements met! \n\n"
		s += ("\nPress any key to exit\n")
		return indent.String(s, 1)
	}

	s += ("\nPress any key to exit\n")

	return indent.String(s, 1)
}
