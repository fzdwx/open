package cmd

import (
	"errors"
	"github.com/fzdwx/open/alias"
	"github.com/fzdwx/open/api"
	"github.com/spf13/cobra"
)

func initDel() *cobra.Command {
	var setCmd = &cobra.Command{
		Use:   "del <alias>",
		Short: "Delete custom quick access",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				api.Check(errors.New("syntax error"))
			}

			alias.Del(args[0])
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			api.Check(alias.TruncateFile())
			api.Check(alias.Save())
		},
	}

	return setCmd
}
