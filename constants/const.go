package constants

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	Enter key.Binding
	Back  key.Binding
	Quit  key.Binding
	Focus key.Binding
}

var Keymap = keymap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Focus: key.NewBinding(
		key.WithKeys("tab", "shift+tab", "enter", "up", "down"),
		key.WithHelp("tab", "change focus"),
	),
}
