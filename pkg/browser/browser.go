package browser

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh/pkg/browser"
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
func Open(url string) error {
	err := b.Browse(url)

	if err == nil {
		fmt.Printf("%s %s %s",
			lipgloss.NewStyle().Bold(true).Foreground(highlight).Render("√"),
			"open",
			lipgloss.NewStyle().Bold(true).Foreground(special).Render(url),
		)
	}

	return err
}
