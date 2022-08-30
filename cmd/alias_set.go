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

func initSet() *cobra.Command {

	var setCmd = &cobra.Command{
		Use:   "set <alias> <PATH>",
		Short: "Customize shortcuts for some paths to alias",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				api.Check(errors.New("syntax error"))
			}

			alias.Set(args[0], args[1])
			api.Check(alias.Save())

			fmt.Println(
				lipgloss.NewStyle().Foreground(lipgloss.Color("118")).Render(symbol.Yes),
				lipgloss.NewStyle().Bold(true).Render("Aliased"),
				lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true).Render(args[0]),
				lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Bold(true).Render(symbol.To),
				lipgloss.NewStyle().Underline(true).Foreground(lipgloss.Color("254")).Render(args[1]))
		},
	}

	return setCmd
}
