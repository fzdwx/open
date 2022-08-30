package cmd

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/open/alias"
	"github.com/fzdwx/open/api"
	"github.com/fzdwx/open/symbol"
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

			if err := alias.TruncateFile(); err != nil {
				api.Eprintln(err)
				return
			}

			api.Check(alias.Save())
			fmt.Println(
				lipgloss.NewStyle().Foreground(lipgloss.Color("118")).Render(symbol.Yes),
				lipgloss.NewStyle().Bold(true).Render("Removed alias"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true).Render(args[0]))
		},
	}

	return setCmd
}
