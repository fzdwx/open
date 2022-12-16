package url

import (
	"github.com/fzdwx/open/internal/browser"
	"github.com/spf13/cobra"
)

var (
	urlCmd = &cobra.Command{
		Use:     "url",
		Short:   "open the specified url",
		Example: `$ open url http://localhost:1313/`,
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(browser.Open(args[0]))
		},
	}
)

func Command() *cobra.Command {
	return urlCmd
}
