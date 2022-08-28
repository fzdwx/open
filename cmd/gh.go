package cmd

import (
	"github.com/fzdwx/open/api"
	"github.com/spf13/cobra"
)

const (
	baseUrl = "https://github.com/"
)

var profile bool

// ghCmd represents the gh command
var ghCmd = &cobra.Command{
	Use:   "gh",
	Short: "Open github in browser",
	Run: func(cmd *cobra.Command, args []string) {
		url := baseUrl
		if profile {
			url = baseUrl + api.UserName()
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
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ghCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
