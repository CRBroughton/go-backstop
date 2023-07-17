package resultsTable

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/config"
	"github.com/crbroughton/go-backstop/constants"
	"github.com/crbroughton/go-backstop/styles"
	"github.com/crbroughton/go-backstop/utils"
)

type GoBackToSettingsMenu bool

type Model struct {
	table table.Model
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Back):
			return m, func() tea.Msg {
				return GoBackToSettingsMenu(true)
			}
		case key.Matches(msg, constants.Keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Enter):
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return styles.TitleStyle.Render(" BackstopJS Test Results: ") + "\n" + styles.ResultsTableStyle.Render(m.table.View()) + "\n" + " Press ESC to return to main menu"
}

func New() Model {
	columns := []table.Column{
		{Title: "Label", Width: 25},
		{Title: "Viewport", Width: 25},
		{Title: "status", Width: 10},
	}

	rows, err := config.GetTestResults()

	if utils.IsError(err) {
		log.Fatal("Could not get test results")
	}

	width, height := config.GetTableWidthHeight()

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithWidth(width),
		table.WithHeight(height),
	)
	t.SetStyles(styles.CreateTableStyles())

	m := Model{t}
	return m
}
