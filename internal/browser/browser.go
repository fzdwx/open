package browser

import (
	"bufio"
	"fmt"
	gitbBrowser "github.com/cli/browser"
	"github.com/fzdwx/open/internal/util"
	"github.com/gookit/goutil/strutil"
	"io"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh/pkg/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/history"
	"github.com/gookit/slog"
	"golang.design/x/clipboard"
)

var (
	b browser.Browser
)

func init() {
	gitbBrowser.Stdout = io.Discard
	gitbBrowser.Stderr = io.Discard
	b = browser.New("", gitbBrowser.Stdout, gitbBrowser.Stderr)
}

// Open url in browser
//
// 1. print url
// 2. append to history
// 3. open in browser
func Open(url string) error {
	ch := make(chan error)
	go func() {
		ch <- open(url)
	}()

	tiktok := time.After(1 * time.Second)

	f := func() error {
		fmt.Printf("%s %s %s\n",
			lipgloss.NewStyle().Bold(true).Foreground(util.Highlight).Render("√"),
			"open",
			lipgloss.NewStyle().Bold(true).Foreground(util.Special).Render(url),
		)

		if err := history.Write(url); err != nil {
			return err
		}

		return nil
	}

	select {
	case err := <-ch:
		if err == nil {
			return f()
		}

		return err
	case <-tiktok:
		return f()
	}
}

func open(url string) error {
	if IsLocalFilePath(url) {
		return nil
	}

	return b.Browse(url)
}

// OpenFromClipboard read url from clipboard and open it.
func OpenFromClipboard() error {
	read := clipboard.Read(clipboard.FmtText)
	if len(read) == 0 {
		return cons.ClipboardEmptyError
	}

	url := string(read)

	if IsNotWebUrlOrLocalFilePath(url) {
		return cons.PathIsNotValidError
	}

	slog.Debug("read url from clipboard: %s", url)

	return Open(url)
}

// OpenFromStdin read url from stdin
//
// ignore more than 4096
func OpenFromStdin() error {

	reader := bufio.NewReader(os.Stdin)

	bytes := make([]byte, 4096)
	n, err := reader.Read(bytes)

	if n == 0 {
		return cons.ClipboardEmptyError
	}

	if err != nil {
		return err
	}
	return Open(string(bytes))
}

// IsWebUrlOrLocalFilePath check s is web url or local file path?
func IsWebUrlOrLocalFilePath(s string) bool {
	if strutil.IsStartsOf(s, []string{cons.HttpPrefix, cons.HttpsPrefix}) {
		return true
	}

	_, err := os.Open(s)

	return err == nil
}

// IsNotWebUrlOrLocalFilePath check s is not web url or local file path?
func IsNotWebUrlOrLocalFilePath(s string) bool {
	return !IsWebUrlOrLocalFilePath(s)
}

// IsLocalFilePath check s is local file path?
func IsLocalFilePath(s string) bool {
	_, err := os.Open(s)

	return err == nil
}
