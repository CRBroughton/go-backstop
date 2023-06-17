package first

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/crbroughton/go-backstop/iterator"
)

type item struct {
	title string
	desc  string
	ID    iterator.Page
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func Content() []list.Item {
	return []list.Item{
		item{title: "Go to settings", desc: "This will take you to the settings", ID: iterator.SettingsPage},
		item{title: "Go to settings", desc: "This will take you to the settings", ID: iterator.SettingsPage},
	}
}
