package gh

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/singleselect"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/open/internal/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/user"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/slog"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"regexp"
)

var repo = &cobra.Command{
	Use:   "repo",
	Short: "open github repository in browser. eg: https://github.com/fzdwx/open",
	Example: `$ open gh repo # must in git project
$ open gh repo fzdwx/open # open https://github.com/fzdwx/open
$ open gh repo open # open # https://github.com/{username}/open`,
	Aliases: []string{"."},
	Run: func(cmd *cobra.Command, args []string) {
		var url string

		if len(args) > 0 {
			url = getUrlFromArgs(args)
		} else {
			url = getUrlInGitProject()
		}

		cobra.CheckErr(browser.Open(url))
	},
}

var (
	urlMatch = regexp.MustCompile("(http|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")
)

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

// get url from ssh git url
// parse from:
//
//	origin  git@github.com:fzdwx/open.git (fetch)
//	origin  git@github.com:fzdwx/open.git (push)
func getFromSshGitUrl(url string) string {
	// 匹配 GitHub 仓库 URL 的正则表达式
	r := regexp.MustCompile(`git@github\.com:(\S+).git`)

	// 进行匹配
	m := r.FindStringSubmatch(url)
	if m == nil {
		return ""
	}
	repoPath := cons.GithubUrl + "/" + m[1]
	return repoPath
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
