package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/crbroughton/go-backstop/iterator"
	"github.com/crbroughton/go-backstop/styles"
	"github.com/crbroughton/go-backstop/utils"
	"github.com/crbroughton/go-backstop/views/settings"
)

const divisor = 4

type item struct {
	title string
	desc  string
	ID    iterator.Page
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	focused iterator.Page
	lists   []list.Model
	loaded  bool
}

func New() *Model {
	return &Model{}
}

func Content() []list.Item {
	return []list.Item{
		item{title: "Run tests", desc: "Runs all stored tests "},
		item{title: "Create new test", desc: "Create a new test"},
		item{title: "Settings Page", desc: "Update your personal settings", ID: iterator.SettingsPage},
	}
}

func (model *Model) initLists(width, height int) {
	delegate := list.NewDefaultDelegate()

	delegate.Styles.SelectedTitle = styles.SelectedItem
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedTitle.Copy()

	defaultList := list.New(
		[]list.Item{},
		delegate,
		width/divisor, // this probably needs fixing
		height/2,
	)
	defaultList.SetShowHelp(false) // without this, the page borks

	model.lists = []list.Model{
		defaultList,
		defaultList,
	}

	model.lists[iterator.MainPage].Title = "Go, Backstop!"
	model.lists[iterator.SettingsPage].Title = "Settings"

	model.lists[iterator.MainPage].Styles.Title = styles.TitleStyle
	model.lists[iterator.SettingsPage].Styles.Title = styles.TitleStyle

	model.lists[iterator.MainPage].SetItems(Content())
	model.lists[iterator.SettingsPage].SetItems(settings.Content())

	model.lists[iterator.SettingsPage].SetSize(200, 20)

}

func (model Model) Init() tea.Cmd {
	return nil
}

func (model *Model) setView(id iterator.Page) {
	switch id {
	case iterator.MainPage:
		model.focused = iterator.MainPage
	case iterator.SettingsPage:
		model.focused = iterator.SettingsPage
	}
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		if !model.loaded {
			model.initLists(msg.Width, msg.Height)
			model.loaded = true
		}
		h, v := styles.AppStyle.GetFrameSize()
		model.lists[model.focused].SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			item, ok := model.lists[model.focused].SelectedItem().(item)
			if !ok {
				model.focused = 0
			}
			if ok {
				model.setView(item.ID)
			}
		}
	}

	var cmd tea.Cmd
	model.lists[model.focused], cmd = model.lists[model.focused].Update(msg)
	return model, cmd
}

func (model Model) View() string {
	if model.loaded {
		mainView := model.lists[iterator.MainPage].View()
		settingsView := model.lists[iterator.SettingsPage].View()

		switch model.focused {
		case iterator.MainPage:
			return styles.AppStyle.Render(mainView)
		case iterator.SettingsPage:
			return styles.AppStyle.Render(settingsView)
		default:
			return styles.AppStyle.Render(mainView)
		}
	} else {
		return "loading..."
	}

}

func main() {
	model := New()
	program := tea.NewProgram(model, tea.WithAltScreen())

	_, err := program.Run()

	if utils.IsError(err) {
		fmt.Println(err)
		os.Exit(1)
	}
}
