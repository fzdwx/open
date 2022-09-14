package browser

import (
	"github.com/cli/go-gh/pkg/browser"
	"os"
)

var (
	b browser.Browser
)

func init() {
	b = browser.New("", os.Stdout, os.Stderr)
}

// Open url in browser
func Open(url string) error {
	return b.Browse(url)
}
