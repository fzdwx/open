package gh

import (
	"github.com/fzdwx/open/internal/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/user"
	"github.com/gookit/goutil/strutil"
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

// get url in git project
func getUrlInGitProject() string {
	buffer := &strutil.Buffer{}
	command := exec.Command("git", "remote", "-v")
	command.Stdout = buffer
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	cobra.CheckErr(command.Run())
	remotePairs := strutil.Split(buffer.String(), "\n")

	pattern, err := regexp.Compile("(http|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")
	if err != nil {
		cobra.CheckErr(err)
	}

	return pattern.FindString(remotePairs[0])
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
