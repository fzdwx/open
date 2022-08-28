package api

import (
	"github.com/cli/go-gh/pkg/browser"
	"io"
)

var (
	b browser.Browser
)

func InitBrowser(writer io.Writer) {
	b = browser.New("", writer, writer)
}

func Browse(url string) error {
	err := Append(url)
	if err != nil {
		return err
	}

	return b.Browse(url)
}
