package history

import (
	"github.com/fzdwx/open/internal/cons"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var (
	historyCmd = &cobra.Command{
		Use:     "history",
		Short:   "Show open history",
		Example: `$ open history`,
		Run: func(cmd *cobra.Command, args []string) {
			command := exec.Command(cons.PreviewCommand(), cons.HistoryFileName())
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			command.Stdin = os.Stdin

			if err := command.Run(); err != nil {
				cmd.PrintErrln(err)
			}
		},
	}
)

func Command() *cobra.Command {
	return historyCmd
}
