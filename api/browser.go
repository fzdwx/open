package api

import (
	"github.com/cli/go-gh/pkg/browser"
	"github.com/fzdwx/open/history"
	"io"
)

var (
	b browser.Browser
)

func InitBrowser(writer io.Writer) {
	b = browser.New("", writer, writer)
}

func Browse(url string) error {
	err := history.Append(url)
	if err != nil {
		return err
	}
	return b.Browse(url)
}

func BrowseWithCheck(url string) {
	err := history.Append(url)
	Check(err)
	Check(b.Browse(url))
}
