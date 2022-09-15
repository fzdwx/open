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
		Use:   "gh [subcommand]",
		Short: "open github in browser",
		Example: `$ open gh -> open https://github.com
$ open gh -s fzdwx -> open https://github.com/search?q=fzdwx`,
		Run: func(cmd *cobra.Command, args []string) {
			url := cons.GithubUrl

			if strx.IsNotBlank(ghSearchString) {
				url = fmt.Sprintf("%s/search?q=%s", url, strx.URLEncode(ghSearchString))
			}

			cobra.CheckErr(browser.Open(url))
		},
	}

	ghSearchString = ""
)

func Command() *cobra.Command {
	gh.AddCommand(profile)

	gh.Flags().StringVarP(&ghSearchString, "search", "s", ghSearchString, "search keyword in github. eg: https://github.com/search?q=fzdwx")

	return gh
}
