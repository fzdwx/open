package log

import (
	"github.com/fzdwx/open/internal/cons"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var (
	cmd = &cobra.Command{
		Use:     "log",
		Short:   "Show open log",
		Example: `$ open log`,
		Run: func(cmd *cobra.Command, args []string) {
			command := exec.Command(cons.PreviewCommand(), cons.LogFileName())
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
	return cmd
}
