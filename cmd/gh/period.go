package gh

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/singleselect"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/open/internal/browser"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/slog"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"regexp"
)

var period = &cobra.Command{
	Use:     ".",
	Short:   "open current github repository in browser",
	Example: `$ open gh . # must in git project`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(browser.Open(getUrlInGitProject()))
	},
}

// get url in git project
func getUrlInGitProject() string {
	buffer := &strutil.Buffer{}
	command := exec.Command("git", "remote", "-v")
	command.Stdout = buffer
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	cobra.CheckErr(command.Run())
	remotePairs := strutil.Split(buffer.String(), "\n")
	slog.Infof("remote pairs: %v", remotePairs)

	if len(remotePairs) < 1 {
		cobra.CheckErr("can not get remote url")
	}

	m := map[string]bool{}
	// get url from ssh git url
	for _, v := range remotePairs {
		url := getFromSshGitUrl(v)
		if strutil.IsNotBlank(url) {
			m[url] = true
			continue
		}
		url = urlMatch.FindString(v)
		if strutil.IsNotBlank(url) {
			m[url] = true
		}

	}

	if len(m) < 1 {
		cobra.CheckErr("can not get remote url")
	}

	var urls []string
	for url := range m {
		urls = append(urls, url)
	}
	if len(urls) == 1 {
		return urls[0]
	}

	selectKeymap := singleselect.DefaultSingleKeyMap()
	selectKeymap.Confirm = key.NewBinding(
		key.WithKeys("enter"),
	)
	selectKeymap.Choice = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select remote url"),
	)

	input := components.NewInput()
	input.Prompt = "FILTER: "
	input.PromptStyle = style.New().Fg(lipgloss.Color("#34d399")).Italic()

	display, err := infinite.NewSingleSelect(
		urls,
		singleselect.WithKeyBinding(selectKeymap),
		singleselect.WithFilterInput(input),
	).Display("Select a remote url")
	if err != nil {
		cobra.CheckErr(err)
	}
	return urls[display]
}

var (
	urlMatch     = regexp.MustCompile("(http|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")
	gitUrlMatch  = regexp.MustCompile(`git@([a-zA-Z0-9.-]+):`)
	gitRepoMatch = regexp.MustCompile(`:([a-zA-Z0-9/._-]+).git`)
)

// get url from ssh git url
// parse from:
//
//	origin  git@github.com:fzdwx/open.git (fetch)
//	origin  git@github.com:fzdwx/open.git (push)
func getFromSshGitUrl(url string) string {
	matches1 := gitUrlMatch.FindStringSubmatch(url)
	if len(matches1) == 0 {
		return ""
	}
	matches2 := gitRepoMatch.FindStringSubmatch(url)
	if len(matches2) == 0 {
		return ""
	}
	return fmt.Sprintf("https://%s/%s", matches1[1], matches2[1])
}
