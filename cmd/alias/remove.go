package alias

import "github.com/spf13/cobra"

var (
	remove = &cobra.Command{
		Use:     "remove",
		Short:   "Remove custom aliases",
		Aliases: []string{"rm"},
		Example: `$ open alias remove blog`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)
