package gh

import (
	"github.com/fzdwx/open/internal/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/user"
	"github.com/gookit/goutil/strutil"
	"github.com/spf13/cobra"
)

var repo = &cobra.Command{
	Use:   "repo",
	Short: "open github repository in browser. eg: https://github.com/fzdwx/open",
	Example: `$ open gh repo # must in git project
$ open gh repo fzdwx/open # open https://github.com/fzdwx/open
$ open gh repo open # open # https://github.com/{username}/open`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(browser.Open(getUrlFromArgs(args)))
	},
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
