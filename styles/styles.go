package styles

import "github.com/charmbracelet/lipgloss"

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
)
