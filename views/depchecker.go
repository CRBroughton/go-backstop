package depchecker

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/crbroughton/go-backstop/config"
	"github.com/crbroughton/go-backstop/utils"
)

type errMsg error

type Model struct {
	Spinner  spinner.Model
	quitting bool
	err      error
	docker   bool
}

func New() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return Model{
		Spinner: s,
		docker:  true,
	}
}

func (m Model) Init() tea.Cmd {
	m.checkDocker()
	return m.Spinner.Tick
}

func (m *Model) checkDocker() {
	_, err := exec.LookPath("docker")

	if utils.IsError(err) {
		m.docker = false
		return
		// fmt.Println("Could not find docker; Please install docker")
		// os.Exit(1)
	}

	config.CreateJSON()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	return fmt.Sprintf("\n\n   %s Checking dependencies, press ESC to quit...\n\n", m.Spinner.View())
}
