package gh

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/singleselect"
	"github.com/fzdwx/iter"
	"github.com/fzdwx/iter/stream"
	"github.com/fzdwx/open/internal/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/user"
	"github.com/google/go-github/v52/github"
	"github.com/gookit/goutil/strutil"
	"github.com/spf13/cobra"
)

var repo = &cobra.Command{
	Use:   "repo",
	Short: "open github repository in browser. eg: https://github.com/fzdwx/open",
	Example: `$ open gh repo # interactive search
$ open gh repo fzdwx/open # open https://github.com/fzdwx/open
$ open gh repo open # open # https://github.com/{username}/open`,
	Run: func(cmd *cobra.Command, args []string) {
		var url string

		if len(args) == 0 {
			url = interactiveSearch()
		} else {
			url = getUrlFromArgs(args)
		}

		cobra.CheckErr(browser.Open(url))
	},
}

func interactiveSearch() string {
	m := &model{}
	_, err := tea.NewProgram(m).Run()
	if err != nil {
		return ""
	}

	res := m.Res()
	if res == "" {
		return ""
	}

	return cons.GithubUrl + "/" + res
}

type model struct {
	selection   *components.Selection
	filterInput *components.Input
	items       []string
	client      *github.Client
	nextPage    int
}

func (m *model) Init() tea.Cmd {
	m.client = github.NewClient(nil)
	m.filterInput = components.NewInput()
	m.selection = components.NewSelection(m.items)
	m.selection.RowRender = func(cursorSymbol string, hintSymbol string, choice string) string {
		return fmt.Sprintf("%s %s", cursorSymbol, choice)
	}
	keyMap := singleselect.DefaultSingleKeyMap()
	m.selection.Keymap = components.SelectionKeyMap{
		Up:   keyMap.Up,
		Down: keyMap.Down,
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "finish select"),
		),
		Choice: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "finish select"),
		),
		Quit:     keyMap.Quit,
		NextPage: keyMap.NextPage,
		PrevPage: keyMap.PrevPage,
	}
	m.selection.FilterInput = m.filterInput

	return m.selection.Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.nextPage = 0
			m.refreshItem()
			m.selection.Choices = mapper(m.items)
			m.selection.RefreshChoices()
			return m, nil
		case "esc":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	_, cmd = m.selection.Update(msg)
	return m, cmd
}

func mapper(items []string) []components.SelectionItem {
	idx := 0
	return stream.Map[string, components.SelectionItem](iter.Stream(items), func(item string) components.SelectionItem {
		selectionItem := components.SelectionItem{
			Val: item,
			Idx: idx,
		}
		idx++
		return selectionItem
	}).ToArray()
}

func (m *model) View() string {
	return m.selection.View()
}

func (m *model) refreshItem() {
	value := m.filterInput.Value()
	repositories, resp, err := m.client.Search.Repositories(context.Background(), value, &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page: m.nextPage,
		},
	})
	if err != nil {
		return
	}
	m.nextPage = resp.NextPage
	m.items = iter.Stream(repositories.Repositories).MapTo(func(repository *github.Repository) string {
		return repository.GetFullName()
	}).ToArray()
}

func (m *model) Res() string {
	value := m.selection.Value()
	if len(value) == 0 {
		return ""
	}
	return m.items[value[0]]
}

// get url from args
// args: [fzdwx/open] => https://github.com/fzdwx/open
// args [open] => https://github.com/{username}/open
func getUrlFromArgs(args []string) string {
	url := cons.GithubUrl

	paris := strutil.Split(args[0], "/")
	if len(paris) > 1 {
		url = url + "/" + args[0]
	} else {
		if user.Name() == "" {
			cobra.CheckErr("can not get user name")
		}
		url = url + "/" + user.Name() + "/" + paris[0]
	}
	return url
}
