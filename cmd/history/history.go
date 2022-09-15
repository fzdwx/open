package history

import "github.com/spf13/cobra"

var (
	historyCmd = &cobra.Command{
		Use:     "history",
		Short:   "Show open history",
		Example: `$ open history`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Command() *cobra.Command {
	return historyCmd
}
