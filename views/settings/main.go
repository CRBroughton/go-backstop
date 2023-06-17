package settings

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/crbroughton/go-backstop/iterator"
)

type Task struct {
	title string
	desc  string
	ID    iterator.Page
}

func (i Task) Title() string       { return i.title }
func (i Task) Description() string { return i.desc }
func (i Task) FilterValue() string { return i.title }

func Content() []list.Item {
	return []list.Item{
		Task{title: "Create user cookie", desc: "test"},
		Task{title: "Create user cookie", desc: "test"},
	}
}
