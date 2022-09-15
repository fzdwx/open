package gh

import (
	"fmt"
	"github.com/fzdwx/open/pkg/browser"
	"github.com/fzdwx/open/pkg/cons"
	"github.com/fzdwx/x/strx"
	"github.com/spf13/cobra"
)

var (
	gh = &cobra.Command{
		Use:   "gh [search keyword] | [subcommand]",
		Short: "open github in browser",
		Example: `$ open gh        -> open https://github.com
$ open gh fzdwx  -> open https://github.com/search?q=fzdwx`,
		Run: func(cmd *cobra.Command, args []string) {
			url := cons.GithubUrl

			if len(args) >= 1 && strx.IsNotBlank(args[0]) {
				url = fmt.Sprintf("%s/search?q=%s", url, strx.URLEncode(args[0]))
			}

			cobra.CheckErr(browser.Open(url))
		},
	}
)

func Command() *cobra.Command {
	gh.AddCommand(profile)

	return gh
}
