package gh

import (
	"fmt"
	"github.com/fzdwx/open/internal/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/user"
	"github.com/spf13/cobra"
)

var profile = &cobra.Command{
	Use:     "profile",
	Aliases: []string{"p"},
	Short:   "open your github profile in browser. eg: https://github.com/fzdwx",
	Example: `$ open gh profile
$ open gh p`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(browser.Open(fmt.Sprintf("%s/%s", cons.GithubUrl, user.Name())))
	},
}
