package cmd

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/fzdwx/open/api"
	"github.com/spf13/cobra"
)

const (
	baseUrl   = "https://github.com/"
	searchUrl = "https://github.com/search?q="
	starUrl   = "https://github.com/%s?tab=stars&q="
)

var (
	profile bool
	search  bool
	star    bool
)

// ghCmd represents the gh command
var ghCmd = &cobra.Command{
	Use:   "gh",
	Short: "Open github in browser",
	Example: heredoc.Doc(
		`$ open gh  -p
		$ open gh  fzdwx
		$ open gh  fzdwx/open
	`),
	Run: func(cmd *cobra.Command, args []string) {
		url := baseUrl
		if profile {
			url = baseUrl + api.UserName()
		} else if search {
			url = searchUrl + args[0]
		} else if star {
			url = fmt.Sprintf(starUrl, api.UserName())

			if len(args) > 0 {
				url += args[0]
			}

		} else {
			if len(args) > 0 {
				url = baseUrl + args[0]
			}
		}
		api.Check(api.Browse(url))
	},
}

func init() {
	rootCmd.AddCommand(ghCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ghCmd.PersistentFlags().String("foo", "", "A help for foo")

	ghCmd.Flags().BoolVarP(&profile, "profile", "p", false, "Open your github profile(auth in gh/cli) in browser")
	ghCmd.Flags().BoolVarP(&search, "search", "s", false, "Use github search you input keywords")
	ghCmd.Flags().BoolVar(&star, "star", false, "Open your github stars tab")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ghCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
