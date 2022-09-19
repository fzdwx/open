package browser

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/fzdwx/open/pkg/cons"
	"github.com/fzdwx/open/pkg/history"
	"github.com/gookit/slog"
	"golang.design/x/clipboard"
	"os"
)

var (
	b         browser.Browser
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
)

func init() {
	b = browser.New("", os.Stdout, os.Stderr)
}

// Open url in browser
//
// 1. print url
// 2. append to history
// 3. open in browser
func Open(url string) error {
	err := b.Browse(url)

	if err == nil {
		fmt.Printf("%s %s %s\n",
			lipgloss.NewStyle().Bold(true).Foreground(highlight).Render("√"),
			"open",
			lipgloss.NewStyle().Bold(true).Foreground(special).Render(url),
		)

		if err := history.Write(url); err != nil {
			return err
		}
	}

	return err
}

// OpenFromClipboard read url from clipboard and open it.
func OpenFromClipboard() error {
	read := clipboard.Read(clipboard.FmtText)
	if len(read) == 0 {
		return cons.ClipboardEmptyError
	}

	url := string(read)

	slog.Debug("read url from clipboard: %s", url)

	return Open(url)
}
