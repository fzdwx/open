package alias

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	as "github.com/fzdwx/open/internal/alias"
	"github.com/fzdwx/open/internal/util"
	"github.com/spf13/cobra"
)

var (
	remove = &cobra.Command{
		Use:     "remove",
		Short:   "Remove custom aliases",
		Aliases: []string{"rm"},
		Example: `$ open alias remove blog`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("url is required")
			}

			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			val, err := as.Remove(args[0])
			cobra.CheckErr(err)

			fmt.Printf("%s %s %s %s\n",
				lipgloss.NewStyle().Bold(true).Foreground(util.Highlight).Render("√"),
				lipgloss.NewStyle().Bold(true).Foreground(util.Special).Render(fmt.Sprintf("%s", val.Name)),
				lipgloss.NewStyle().Bold(true).Foreground(util.Red).Render("≠"),
				val.Url,
			)
		},
	}
)
