package styles

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	AppStyle = lipgloss.NewStyle().Padding(1, 2).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true)

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#5853e5")).
			Padding(1, 2).
			Bold(true)

	SelectedItem = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#39d884")).
			BorderLeftForeground(lipgloss.Color("#39d884")).
			Padding(0, 2)

	DocStyle = lipgloss.NewStyle().Margin(1, 2)

	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	MenuItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	HighlightedMenuItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	ResultsTableStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240"))
)

func CreateTableStyles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	return s
}
