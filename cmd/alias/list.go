package alias

import "github.com/spf13/cobra"

var (
	list = &cobra.Command{
		Use:     "list",
		Short:   "Show custom aliases",
		Aliases: []string{"ls"},
		Example: `$ open alias list`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)
