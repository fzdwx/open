package cmd

import (
	"errors"
	"github.com/fzdwx/open/alias"
	"github.com/fzdwx/open/api"
	"github.com/spf13/cobra"
)

func initSet() *cobra.Command {
	var setCmd = &cobra.Command{
		Use:   "set <alias> <PATH>",
		Short: "Customize shortcuts for some paths to alias",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				api.Check(errors.New("syntax error"))
			}

			alias.Set(args[0], args[1])
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			api.Check(alias.Save())
		},
	}

	return setCmd
}
